package fwbundle

import (
	"archive/zip"
	"bytes"
	"compress/flate"
	"encoding/json"
	"io"
	"io/ioutil"
	"path"
	"path/filepath"

	"cesanta.com/common/go/ourutil"
	"github.com/cesanta/errors"
	"github.com/golang/glog"
)

const (
	ManifestFileName = "manifest.json"
)

func ReadZipFirmwareBundle(fname string) (*FirmwareBundle, error) {
	var r *zip.Reader

	zipData, err := ourutil.ReadOrFetchFile(fname)
	if err != nil {
		return nil, errors.Trace(err)
	}

	r, err = zip.NewReader(bytes.NewReader(zipData), int64(len(zipData)))
	if err != nil {
		return nil, errors.Annotatef(err, "%s: invalid firmware file", fname)
	}

	fwb := NewBundle()

	blobs := make(map[string][]byte)

	for _, f := range r.File {
		rc, err := f.Open()
		if err != nil {
			return nil, errors.Annotatef(err, "%s: failed to open", fname)
		}
		data, err := ioutil.ReadAll(rc)
		if err != nil {
			return nil, errors.Annotatef(err, "%s: failed to read", fname)
		}
		rc.Close()
		blobs[path.Base(f.Name)] = data
	}
	manifestData := blobs[ManifestFileName]
	if manifestData == nil {
		return nil, errors.Errorf("%s: no %s in the archive", fname, ManifestFileName)
	}
	err = json.Unmarshal(manifestData, &fwb.FirmwareManifest)
	if err != nil {
		return nil, errors.Annotatef(err, "%s: failed to parse manifest", fname)
	}
	for n, p := range fwb.FirmwareManifest.Parts {
		p.Name = n
		p.SetDataProvider(func(name, src string) ([]byte, error) {
			data, ok := blobs[src]
			if !ok {
				return nil, errors.Errorf("%s not found in the archive", src)
			}
			return data, nil
		})
	}
	return fwb, nil
}

func WriteZipFirmwareBundle(fwb *FirmwareBundle, fname string, compress bool) error {
	buf := new(bytes.Buffer)
	zw := zip.NewWriter(buf)
	zw.RegisterCompressor(zip.Deflate, func(out io.Writer) (io.WriteCloser, error) {
		if compress {
			return flate.NewWriter(out, flate.BestCompression)
		} else {
			return flate.NewWriter(out, flate.NoCompression)
		}
	})
	// Rewrite sources to be relative to archive.
	for _, p := range fwb.Parts {
		if p.Src == "" {
			continue
		}
		data, err := p.GetData()
		if err != nil {
			return errors.Annotatef(err, "%s: failed to calculate checksum", p.Name)
		}
		p.SetData(data)
		p.Src = filepath.Base(p.Src)
		if err := p.CalcChecksum(); err != nil {
			return errors.Annotatef(err, "%s: failed to calculate checksum", p.Name)
		}
	}
	manifestData, err := json.MarshalIndent(&fwb.FirmwareManifest, "", " ")
	if err != nil {
		return errors.Annotatef(err, "error marshalin manifest")
	}
	glog.V(1).Infof("Manifest:\n%s", string(manifestData))
	zf, err := zw.Create(ManifestFileName)
	if err != nil {
		return errors.Annotatef(err, "error adding %s", ManifestFileName)
	}
	if _, err := zf.Write(manifestData); err != nil {
		return errors.Annotatef(err, "error adding %s", ManifestFileName)
	}
	for _, p := range fwb.Parts {
		if p.Src == "" {
			continue
		}
		data, err := p.GetData()
		if err != nil {
			return errors.Annotatef(err, "error getting data for %s", p.Name)
		}
		if zf, err = zw.Create(p.Src); err != nil {
			return errors.Annotatef(err, "%s: error adding %s", p.Name, p.Src)
		}
		if _, err = zf.Write(data); err != nil {
			return errors.Annotatef(err, "%s: error adding %s", p.Name, p.Src)
		}
	}
	if err = zw.Close(); err != nil {
		return errors.Annotatef(err, "error closing the archive")
	}
	return ioutil.WriteFile(fname, buf.Bytes(), 0644)
}
