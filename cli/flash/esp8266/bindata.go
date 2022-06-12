// Code generated for package esp8266 by go-bindata DO NOT EDIT. (@generated)
// sources:
// stub/stub.json
package esp8266

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"time"
)
type asset struct {
	bytes []byte
	info  os.FileInfo
}

type bindataFileInfo struct {
	name    string
	size    int64
	mode    os.FileMode
	modTime time.Time
}

// Name return file name
func (fi bindataFileInfo) Name() string {
	return fi.name
}

// Size return file size
func (fi bindataFileInfo) Size() int64 {
	return fi.size
}

// Mode return file mode
func (fi bindataFileInfo) Mode() os.FileMode {
	return fi.mode
}

// Mode return file modify time
func (fi bindataFileInfo) ModTime() time.Time {
	return fi.modTime
}

// IsDir return file whether a directory
func (fi bindataFileInfo) IsDir() bool {
	return fi.mode&os.ModeDir != 0
}

// Sys return file is sys mode
func (fi bindataFileInfo) Sys() interface{} {
	return nil
}

var _stubStubJson = []byte(`{"params_start": 1074790400, "code": "b0150000a073ff3fdf7befbd841042080080000020290000e00d00003b70ff3ff076ff3f801b00000478ff3fa00d000000080000600900000000ffff500e000050120000242900003070ff3f3f70ff3f232900007476ff3ff875ff3f800b00007875ff3ff874ff3f0100000000000000f1ff000048df004090e1004068e2004092a1209011c0e26144ed022804a221482a232261252807bd030c432a8630aa10826128a2612d7cf8a22148026147c26146d26145f2614342613052612b72613182612a9d06378a0b5036c080228020238022612a82212a0c021b68878625a2212ba7391f888ea8ae82612688fe280ea261248261273c53c81ef8ee589e27b31346f60300a22131822130290a29087cd74651040031bcffe042114a333803a003000c03a22148393e392e0c13397e394e30ca1007ea296d09dd0bfd0cc261245d0cc26126863600008221480c128071047cc37032937d03290e6d098d0b86e003fd022261245d02226126cd02a22125a7bbd5220b001b8b292e460600a221480c22207a100c147cc37034937d03290e6d0986d3038d0ba22125a7b8df282e5208004199ff8032115a334033827198ff593e1bd80c143737102c0330751037850820303432c3f830478382212d5d04dce82024418b2252a001718eff0012400035a13737080002406020912058935054206d098ce5460000dd0b2c42290e6d09463e030000f63c39c60500a221480c13a071047cc27023937d020c32c613006d09dd0b2d0da221258d02a7b2dd3208001b22001c400033a18bccdd0230ff20b63ce2f02024295e202141296ef0f341c2ccfd56f21ec02024273c1f860f00a221480c13a071047cc27023937d020c52290ec69b030000006d09dd0b2d0da221258d02a7b2d83208001b22001c400033a18bcc30ff20c03024dd02373cdfc02024000240f0f0917c8220cc100c05168c06b68c1bc60d00a221480c13a071047cc27023937d020c6206e9ff006d09dd0b2d0da221258d02a7b2dc3208001b22001c400033a18bccdd0230ff20b68ce23149ff5a2e3a22f24200c2ccf8f0f841060b008221480c138071047cc27023937d020c724664036d09dd0ba22125a7bde1413dff320d005a2e4a223242001bdd1b55f6450286ddff22de293212115212107cf23022302020f4271573c600006d09dd0b2c7206e902b68c1d460e0000a221480c13a071047cc27023937d023c3246bfff906920b0db202d0da221258d02a7b2db3208001b22001c400033a18bccdd0230ff20b68ce2f08074826126f0f841c2ccf80602003c4206e0026d09dd0ba22128a7b6f08221260b558246001b66cc1546f90256fcf8461d000c9206d7026d09dd0ba22128a7361c46fbff00008221480c138071047cc27023937d022c62462d036d09dd0ba22125a7bde1822128d07ac06028c077b2017d0277b5017d054d073d0d2d0652613862613972613692613ab261370114ffc00000722136622139522138b2213792213a7add7a667055c056f5f8c6d80266320d46010000006d09dd0b0ca206a90226120286200022a12029be21eefe2c0449ce0c532a2e52613862613992613ab261370100ffc0000032a0d0b2213792213a62213952213822ce403a3e42a0084242001b222793f722a0d032a1402a2e3a3e0c944242001b222793f722a14032a1582a2e3a3e0c744242001b222793f722a15832a1602a2e3a3e0c844242001b222793f78662010000000c0521cffe5a22220200273c1d461000a221480c13a071047cc27023937d020cb2c65aff0000006d09dd0b31c4fe2d0d5a33a221258d02a7b2d54208001b22001c400044a140ff204203008bccdd02473cdf21bbfe41bbfe5a223202000c120013400022a1e075110b227a44f022104804000340f0f09130ccc08b35e033113a3e2a2429331b55f6350206ddff21aefe42a1200c032a2e62613992613ab2613701befec0000062213992213ab221370c05c617000000b63c1e860e000000a221480c13a071047cc27023937d020ce2462fff00006d09dd0b2d0da221258d02a7b2db3208001b22001c400033a18bccdd0230ff20b63ce22194fe4193fe5a22220200f030242a2e4a22c2ccfd324200f0f3411b5528de27359c1c3229dec61901004c040c0322c14452613862613992613ab261370199fec000002185fea221294184fe208a8222a1602028800c03202e808261350191fec00000822135217efe42a4802a280c032a2e018cfec00000286e8221352261338b22e022118a8e2a2e0c04783252213862213992213ab2213782612e82612f3d0842612c0606000022033f82212ce022112a214222111b881b4442621182612ca2212c1b33a797df0c0229112901a2c1448d014d020c17381a1b772a23f0221129283a444baa4b8866b7eb315efe3a22164210f62406863f006d09dd0b2c32c69800000022212f42024042612916a40de024112a21880242613c1ba8a90282613d0c02a2213df02211a0a00482213c22613b202a20a2213d0b88a0a14182613ca2613d56c8fd0ca84738347084117088200088118080310c1a826129001440004aa1460400000082212ef0a211aaa88221294a22825ab0a2a3ff27bae9861a004221322080944a88f088118a8e4298b0cc743258b030432032c3fe202941460800a22132aa84f088118a8e82d809429830cc643258304d0332c3fe8221290b888261292021412080048044c07cfa822129404a300cbaa798c722213282a4b04a428a44f044114a4e725400a2212f1b771baaa2612fc6030082212922a6d02028820c077cf3226132a2212c771a0286bcff222133262202069c000684000ce2c7b202463000822125d028c0a62202862500310dfef020943a22f022112a2e22120000221120403196e200202931162205273c020624004612000ca3c7b344a103fe7cf87d03000740f020912020044048304a42aa44f044114a4e4294001b33d644062b7777bcdb06060000a221480c13a071047cc27023937d021c02460c026d09dd0b82212587bde1220d001b3d001c400022a120ff208bcc0ce2dd03c7320246dbff0608000000220d018b3c0013400022a1320d002bdd001c400033a130222020ff20c2cc1031e0fdf020943a22f022112a2e32120000331130403196b30030393140408486090000000081d7fd0ca37cf7000340f020912020044047304a428a44f044114a4e4294001b339624fe4261260cf2000340f0f09130ccc047220e31cbfd5a2e3a224242001b55c63700a2212666ba18dc551c08826126c600006d09dd0b1c12290e7cf746d901000021c1fd4a4242c4f0420400426124473c20c60f00a221480c14a071047cc37034937d031c23390edd02c6cd0100006d09dd0b2d0d8221253d0287b2d7320300a22124001c400033a11b228bccdd0230ff20a73cdf82212431acfda221260c170018400077a1aa330b7732c3f0f02710720300000840f0f0912a7721a1fd80ccc02a2e5a220c0366ba0a41a0fd5a3e4033803203004d0752613862613972613692613ab2613701a3fdc0000052213872213662213992213ab221377a554600000c0548be28ce2a2427b5020678ff579207460200006d09dd0b1c5246c6ff7188fd22ce407a7e3d0752613862613992613ab26137726136018efdc0000038be7221362174fd48ce3a372a2e0189fdc00000b2213792213a622139522138286e0b22296e886e82612996d81446e3fe322125d023c0a6420e4221286024c0a62202464f0086f101c7b802062f00a62202462500f0209432a0b03a22f022112a2e22120000221120403196f200202931163205273c02462400861200000ca3c7b3447cf8a2a4b07d03000740f020912020044048304a42aa44f044114a4e4294001b33d644062b7777bcdb060600008221480c138071047cc27023937d021c724668016d09dd0ba22125a7bde1220d001b3d001c400022a120ff208bcc0ce2dd03c7320206dbff060800220d0132cc080013400022a1320d00d2cd02001c400033a130222020ff20c2cc1032a0b0f020943a22f022112a2e321200003311305031968300303931505084c608000ca37cf742a4b0000340f020912020045057305a524a55f055115a5e5295001b339625fe22a0ff000340f0f09130ccc057a202c64600460200001c8206df006d09dd0b82212887b6f05246001b660ce8a2a10006abffc73819220d01320d00802211302220001c400022a120ff202bddc2cc10f0209432a0b03a22f022112a2e22120000221120703196720020293146090000000ca2000240f030917cf47074303030047a3342a4b04a33f033113a3e7293001b2296d7fd5d07000240f0f09120ccc0a70702062200c73819220d01320d00802211302220001c400022a120ff202bddc2cc1032a0b0f020943a22f022112a2e321200003311304031968300303931060b000000000ca3000340f020912020042261297cf24042302221291b334a4222a4b02a44f044114a4e4294009674fd724600000340f0f09130ccc0a704081b665d0446020000004246012b66066cff50508466f50286b70022aeff2a5521e1fce055115a22280222612421dffc8221245a52580516f805873c1e460f000000a221480c13a071047cc27023937d021c92465afd00006d09dd0b2d0da221258d02a7b2db320800822124001c400033a11b228bccdd0230ff20873cdfa221240c12001a400022a10b22f02210a0ccc0000a40f0f0912055800ce2c7b202c62f00822125d028c0a62202062500f02094f0221122d20f2a2e22120000221120403196e200202931162205273c020624004612000ca3c7b344a1b5fc7cf87d03000740f020912020044048304a42aa44f044114a4e4294001b33d644062b7777bcdb06060000a221480c13a071047cc27023937d021ca246b7006d09dd0b82212587bde1220d001b3d001c400022a120ff208bcc0ce2dd03c73202c6dbff0608000000220d018b3c0013400022a1320d002bdd001c400033a130222020ff20c2cc10f02094f0221122d20f2a2e32120000331130403196930030393140408406090000818afc0ca37cf7000340f020912020044047304a428a44f044114a4e4294001b339624fe000340f0f09130ccc03180fce044114a333803326124317efca221244a434804426126161a06a73c1b860e00a221480c13a071047cc27023937d021cb206f6fc006d09dd0b2d0da221258d02a7b2dc320800822124001c400033a11b228bccdd0230ff20873cdfa221240c12001a400022a18221260b22f022102a88000a40f0f091a0ccc0826126a2212b82212da0a6c0a261279c48822126873a060603006d09dd0b2c52290e8d0dc67700a221278221263d06802ac0a2212a82212ba022102a2827b6013d02a221285a33373a3c461100003c52290e8d0d0c27c66b006d09dd0b82212887b6eaa22127822126802ac0a2212a82212ba022102a28a221272202001baa224600a261271b662d057cf51632c55a5246f2ff3202004d0532460032020152c5fd3246013202023b223246023d063b66e635dfe61502060aff7202007246006625052202012246014a630605ff00285e07e202c694fcc02024273c1b860e00a221480c13a071047cc27023937d022c0206a8fc006d09dd0b2d0da221258d02a7b2dc3208001b22001c400033a18bcc30ff20c03024dd02373cdfc02024000240f0f0917c8220cc100601000bddc2ccf8d7bb02f68cf42116fc3116fcc04c2052613862613992613ab2613705e60082214822c2ff52213862213992213ab22137f0f21007e80286230052a000167c06b68c1d460e0000a221480c13a071047cc27023937d022c924682fc00006d09dd0b2d0da221258d02a7b2db3208001b22001c400033a18bccdd0230ff20b68ce2f03074c2ccf8f0f8418609008221480c138071047cc27023937d022ca2290e8d0d460d006d09dd0ba22125a7bddd320d001bdd284e1b55802211302220294ef64506c6ddff6d09dd0b2c22290e8d0d0c074601006d098d0b7cf72617114b27cc628602000b88c2ccf887bb02f68cf421dcfb31dcfbc91ec04c2052613862613972613682613592613ab2613705d700a22126522138a98ea22124822135a9aea22127b221370b22a9fef0f210a2213062213992213af9ee599eb088c0890a82213190d6c0a22148d9080c92722136278a02c63e0096870f31a8fbf87e2d0d92613a01c4fbc00000f0c0f472213692213af0f0f55d027c880c06462800320900ca433209014a333a44faf44209023a344209033aff3a344209043aff3a344209053aff3a344209063aff3a344209073aff3ac4caff8b992799c28025103d02c60300000020a3c0aaa9a20a001b33aacccaff5733ee3d062735022035c03a9931a2fb2d0c52613862613972613682613592613a01a0fbc00000319bfb20c220f02f20019cfbc0000052213892213a50ddc08221355179fb722136622139fd028c9d502341d022112a2986e2ff00ff11ca4f497edca7822148076815f84e7ce2f044c0407293060200c738020635fe863cfe02214792a120c22146d22145e22144f221432d079a110df0000000f42a0000102b000091ffff12c1f0c92109319011c072c110c2d72a326c3c3d040c0449417c944066100c44526c3d40662041f3ff69015d026d0222c11072d12b4a4285dcfe7cf3cc12322c3c91eeff2d039a110831c82112c1100df01400f03f583f00404446004012c1f002610301fcffc0000001fcffc0000031f8ff0831c0200028030c14402220c02000290312c1100df00014000060ffff0f001ce20040d839004012c1f0d911dd02c0221151f9ffd022c08022115a22c921c02000c80221f6ff02610320cc10302c8240342001f3ffc000003d02d0207401f1ffc0000008312d0cd811c82112c1100df00000000000ff3fff0f000031feff52d34068a57046116a44d044114a33481361faff47360e1b6469134a3322430828b51b2229b50df0000800006000000060081000000c0000601c0000601000006021faff12c1d0c020002802c9a1c1e9fff97109b1d991e9812901f2dc4046400000000c0e21f2ff680cc020002802204074663602062c00f6460a9c2652a0db579425460700664602063900c63200000022a0c02714020630000c12290c69bfc62d00000c32c62b0072a0c058af7794531b252635024600000c0229af48af0c067024114a22d0221141daff2a2c69124a2262420028bf6792237025115a22d022112a2c4a2222020020400467d20e0c42290c21d0ffc020004902861c000c0246150000006616147025115a22d0221151c8ff2a2c5a22424200460e00202074c5edff060d0000000022a0dc27141622a0dd2714180c42290c21bfff0c03c020003902860a0022a0c0060100000022a0db10112085eaff0c22290c1beee71d0246c0ff21b5ffc02000d802d0d074561def380121b2ff08b1c8a1d891e881f871c02000390212c1300df00000005810000070100000741000007810000080100000189800401c4b0040803c0040349800400099004091faff12c1e051f5ffc961d951e941f93109719011c01a55c2d11029052d0cdd03ed0401f2ffc0000021ebff31edff2a411a330c0f4903861f000000f06dc067be016d0e51e6ff4d061a5558053d015a2f51e4ff1a55690501e6ffc000004d0221e1ff1a22680256340451deff3d061a5549052d0101dfffc0000021daff0c431a224802426c1641d6ff1a44280405850066421f422c1647bf1ef044c03d01c02c2001d5ffc00000f22c168602005c32c60900005c42460800d7bf0206dfff21c6ff3d0c1a2201cdffc0000021c3ff1c031a2201c8ffc000000c0291c3ff9a110871c861d851e841f83112c1200df00014c7ff3f04020060000200608c4400403c44004012c1f0c921c1f9ffe901ed02280c0931d911dd0301f9ffc00000280c01f8ffc0000021f3ff0831c02000e90221f2ffc821c02000d902e801d8110c0212c1100df0000000001000000080fe3f30400000b81310402400006064000094ffff000000000100000080000000000100389c1c0040ff3f78480040880f0040a80f00404c4a0040980f004092a1409011c0f2614bfd0222c158c2614e02614fd2614de2614c32613d4261410190ffc00000f020b43c2c56523072213d3c3c7020b456a22f01ecffc00000dd023c4c56d22e3d0241dfff211affc2a0d4015dfac000004142fa21d9ff3d0d015afac0000031d8ff4d0d0c5201e1ffc0000021d6ff31d6ffc020007802cac172613fc020003902211bff32a101c0200039022c0201d8ffc0000022a0b03d0d2c442a210149fac000001c843d0d2d0c0146fac000001c832d0c05680020ea038b7c2261307261428688000000e0ea03060d0072213d31bfffd027c027b31af020f456420131beff2d0fc5e8ff56c21d21baff2addc60300000031b9ff2d0f85e7ff56c21cd2dd107221418c67280c22d210273dbd30ea0322212fe0e2c03aeee2612f50ea0321e3fe71b0ff32d2104d023833281278973a2232d420385361a9ff3a2232d430487372613c2a24291c21d9fe3802460300000040ea035044c047b602465c0071a1ffc0200048a772213c7794066643e2c658000066430206570040ea0332212c5033c04a3332612c70ea0372614072213c42213c7077114a37d0331172613e8b732a773a2231d1fe3a3232030007631ce181ff2d074d0e32a0ff7261430101fac000007221435d0e06080000005812176319e17aff31e1f90c164d072d0e859eff5d02660202863c00e07e2030ea0322212d4221405261444022c03a224d0522612d3d0722c1587261430117ffc0000032213e42213c5221444a2331a5fed022112a2331b2fe3a2222020007e244e0ea03214bff2802014dffc0000020ea0332212f3a22e022c022612fe0ea035221447221434d053d072d0f0166ffc0000052214456b20730ea0322212ee022c03a2222612e72213e32213c418dfe3a27d022117158ff2a240c03391228971b422632054997460100002153ff3992e80c5c845aee4a312d015affe90c01c5f9c000002221423d0101edfec000001c83c02c20454700460800003c5cc613003c6c8612003c7c461100003c8cc60f003c9c860e004c0c460d000000280c72213d77b202c674ff42a0b01a4432c15822c41401dbfec0000020ea033221300c0c3022c022613022a0b02c431a22c54100212aff72213fc0200079022c020133ffc0000002214f92a1402d0cd2214dc2214ee2214cf2214b9a110df012c1f0611bff09315d025c224736042d0545b1ff083112c1100df000b0100000c0100000d010000012c1e091feffc961e941f9310971d951cd039011c0310cfffd02ed049c1422a06247b302062d0021f4ff1a22490286010021f1ff1a223902d1a8fedad12d0d01abfec00000461c0022d11001a8fec0000021e9ffdd0c1a222802c7b20621e6ff1a22d8024d0d3d012d0f01a1fec000008c5222a063c61800002197fe4d0d103120102280019dfec00000ac7e4d0d3d0122d1100199fec0000021d6ff32d1101022800196fec0000021d3ff1c031a220191fec00000daffd0ccc056acf83186fe21cdff3a311a22018dfec0000021c9ff1c031a220188fec000002d0c91c8ff9a110871c861d851e841f83112c1200df00000001040020060ffffff0012c1e00c02290141bdfe21faff026107c961c02000226400c02000380420c3102783f421f5ff31f5ffc02000280230221029010c432d010170fec0000008712d0cc86112c1200df02000006000000200fffffdff1880ff3f8449004012c1d0c9a15c52c1f9ff02610b224115c02000280c31f6ff302220c02000290cc02000280c31f3ff302210c02000290c0c1322c1150525002612024642007cf22241142201150c930b2220207427b302c63a0031e9ffe022112a232802a0020000000032a010102120c5210066a20c482138112801c5abff063000004c12862e001c032d0105200066a20c482138112801c5deff062900005c12862700000032a010102120051e0066a20c48213811280145dfff0621000022a061461f0045eeffc61d0000000001cdffc00000c61a000c022241140c1322c114451900220115c61c0000000032a0101021208519002032202d0146060032a0101021208518006642143801c02000380339413d0222c110c515000607000022af9146070032a01010212005160066820e2221003811c0200039020c0286000022afa122411432a00122c11485120022011532c2fa303074b6230206aeff02210bc2210a12c1300df00000000010401027000050c3000000110040a8100040bc0f00402c4c0040cc2e00404c26004021f7ff12c1e0d951d802e941e8120c020971c961d91101f6ffc00000855bff7158fe5151fe4157fe3159fe62a1000c02c2a00001f0ffc00000c26100c71e1421e8ff01edffc000004d0e3d0d2d0c455cff22610021e4ff01e8ffc0000032a00410212045080005e1ff20c22021ddff01e2ffc00000666c23460100004b2206010031d9ff4b21c0200048023794ed31d7ffc0200039023df08601000001d8ffc000000871c861d851e84112c1200df0003040ff3f0080fe3f2c70ff3f21feff41fcff12c1f02044c040422132a00002610301a8f8c0000021f8ff20122005f2ff083112c1100df00012c1f002610301c9fdc00000083112c1100df000643b004012c1d0d991e981f97109b1c9a1ed02dd03f2a0c001faffc00000cd02f792f40c0f860d0001f6ffc000004d0222a0c027143e22a0db27941b490101f0ffc0000052a0dc480157120952a0dd571205460400004d0cfa2e4242001bffd79fc5c6000000000c0fc2a0c001e5ffc00000c792f608b12d0fc8a1d891e881f87112c1300df00000001440e6c4092033810022a10df000000032a10c020df000", "code_start": 1074790408, "entry": 1074798024, "data": "0203070031302e302e32000505040003030b00756e6b6e6f776e206572726f72006e6f206572726f7200756e646566696e6564206572726f7200746f6f206d616e792066696c65730066696c6520746f6f206c6172676500756e737570706f72746564206d6574686f6400756e737570706f7274656420656e6372797074696f6e00756e737570706f727465642066656174757265006661696c65642066696e64696e672063656e7472616c206469726563746f7279006e6f742061205a4950206172636869766500696e76616c696420686561646572206f72206172636869766520697320636f7272757074656400756e737570706f72746564206d756c74696469736b2061726368697665006465636f6d7072657373696f6e206661696c6564206f72206172636869766520697320636f7272757074656400636f6d7072657373696f6e206661696c656400756e6578706563746564206465636f6d707265737365642073697a65004352432d333220636865636b206661696c656400756e737570706f727465642063656e7472616c206469726563746f72792073697a6500616c6c6f636174696f6e206661696c65640066696c65206f70656e206661696c65640066696c6520637265617465206661696c65640066696c65207772697465206661696c65640066696c652072656164206661696c65640066696c6520636c6f7365206661696c65640066696c65207365656b206661696c65640066696c652073746174206661696c656400696e76616c696420706172616d6574657200696e76616c69642066696c656e616d650062756666657220746f6f20736d616c6c00696e7465726e616c206572726f720066696c65206e6f7420666f756e64006172636869766520697320746f6f206c617267650076616c69646174696f6e206661696c65640077726974652063616c6c65646261636b206661696c65640073747265616d20656e64006e6565642064696374696f6e6172790066696c65206572726f720073747265616d206572726f720064617461206572726f72006f7574206f66206d656d6f727900627566206572726f720076657273696f6e206572726f7200706172616d65746572206572726f7200000000000000000049454e44ae4260820089504e470d0a1a0a0000000d4948445200000000000000000800000000000000000000000049444154000029011040740110409c01104026021040e51010407e021040d20210401c031040e5101040d5031040540410400d051040e5101040e5101040ba051040e51010407c0810403d0910407a091040e5101040e5101040200a1040e51010400c0b1040b50b10400e0d1040d00d10409e0e1040e5101040e5101040e5101040e5101040d60f1040e5101040d6101040d6061040ff011040fc0e1040f803104054031040e51010406e101040b5101040e5101040e5101040e5101040e5101040e5101040e5101040e5101040e510104078031040b10310403b0f10405170ff3f5a70ff3f6a70ff3f7970ff3f8870ff3f9b70ff3fb270ff3fc670ff3fe770ff3ff970ff3f2071ff3f3e71ff3f6b71ff3f7e71ff3f9b71ff3faf71ff3fd271ff3fe471ff3ff571ff3f0872ff3f1a72ff3f2b72ff3f3d72ff3f4e72ff3f5f72ff3f7172ff3f8272ff3f9372ff3fa272ff3fb172ff3fc672ff3fd872ff3f010000000200000003000000040000000500000007000000090000000d00000011000000190000002100000031000000410000006100000081000000c100000001010000810100000102000001030000010400000106000001080000010c00000110000001180000012000000130000001400000016000000000000000000000000000000000000000000000000000000100000001000000020000000200000003000000030000000400000004000000050000000500000006000000060000000700000007000000080000000800000009000000090000000a0000000a0000000b0000000b0000000c0000000c0000000d0000000d0000000000000000000000030000000400000005000000060000000700000008000000090000000a0000000b0000000d0000000f0000001100000013000000170000001b0000001f000000230000002b000000330000003b0000004300000053000000630000007300000083000000a3000000c3000000e3000000020100000000000000000000000000000000000000000000000000000000000000000000000000000000000001000000010000000100000001000000020000000200000002000000020000000300000003000000030000000300000004000000040000000400000004000000050000000500000005000000050000000000000000000000000000000101000001000000040000000000040206000000000000003e70ff3f01000000f072ff3f02000000fb72ff3fffffffff0b73ff3ffeffffff1673ff3ffdffffff2373ff3ffcffffff2e73ff3ffbffffff3c73ff3ffaffffff4673ff3ff0d8ffff5473ff3f000000006410b71dc8206e3bac30d9269041dc76f4516b6b5861b24d3c7105502083b8ed44930ff0e8a3d6d68cb361cbb0c2649bd4d2d38678e20aa01cf2bdbd00000000010000000600000020000000100000002000000080000000000100000002000000030000dc050000000000000100000003000000070000000f0000001f0000003f0000007f000000ff000000ff010000ff030000ff070000ff0f0000ff1f0000ff3f0000ff7f0000ffff000010111200080709060a050b040c030d020e010f00000808090909090a0a0a0a0a0a0a0a0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d00001213141415151616161617171717181818181818181819191919191919191a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1a1b1b1b1b1b1b1b1b1b1b1b1b1b1b1b1b1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1c1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d1d000000000101010102020202020202020303030303030303030303030303030304040404040404040404040404040404040404040404040404040404040404040505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050506060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060606060707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070707070700010203040405050606060607070707080808080808080809090909090909090a0a0a0a0a0a0a0a0a0a0a0a0a0a0a0a0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0b0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0c0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0d0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0e0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f0f10101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101010101111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111000000000000000001010101010101010202020202020202020202020202020203030303030303030303030303030303030303030303030303030303030303030404040404040404040404040404040404040404040404040404040404040404040404040404040404040404040404040404040404040404040404040404040405050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505050505000001010201030104010501060107010801090109010a010a010b010b010c010c010d010d010d010d010e010e010e010e010f010f010f010f011001100110011001110111011101110111011101110111011201120112011201120112011201120113011301130113011301130113011301140114011401140114011401140114011501150115011501150115011501150115011501150115011501150115011501160116011601160116011601160116011601160116011601160116011601160117011701170117011701170117011701170117011701170117011701170117011801180118011801180118011801180118011801180118011801180118011801190119011901190119011901190119011901190119011901190119011901190119011901190119011901190119011901190119011901190119011901190119011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011a011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011b011c011c011c011c011c011c011c011c011c011c011c011c011c011c011c011c011c011c011c011c011c011c011c011c011c011c011c011c011c011c011c011d01e71b1040051c1040231c1040421c10404b1c1040541c1040541c10406a1c10407b1c1040a21c1040", "data_start": 1073705008, "num_params": 2, "code_size": 7840, "data_size": 4112}`)

func stubStubJsonBytes() ([]byte, error) {
	return _stubStubJson, nil
}

func stubStubJson() (*asset, error) {
	bytes, err := stubStubJsonBytes()
	if err != nil {
		return nil, err
	}

	info := bindataFileInfo{name: "stub/stub.json", size: 24084, mode: os.FileMode(420), modTime: time.Unix(1, 0)}
	a := &asset{bytes: bytes, info: info}
	return a, nil
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("Asset %s can't read by error: %v", name, err)
		}
		return a.bytes, nil
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// MustAsset is like Asset but panics when Asset would return an error.
// It simplifies safe initialization of global variables.
func MustAsset(name string) []byte {
	a, err := Asset(name)
	if err != nil {
		panic("asset: Asset(" + name + "): " + err.Error())
	}

	return a
}

// AssetInfo loads and returns the asset info for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func AssetInfo(name string) (os.FileInfo, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		a, err := f()
		if err != nil {
			return nil, fmt.Errorf("AssetInfo %s can't read by error: %v", name, err)
		}
		return a.info, nil
	}
	return nil, fmt.Errorf("AssetInfo %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() (*asset, error){
	"stub/stub.json": stubStubJson,
}

// AssetDir returns the file names below a certain
// directory embedded in the file by go-bindata.
// For example if you run go-bindata on data/... and data contains the
// following hierarchy:
//     data/
//       foo.txt
//       img/
//         a.png
//         b.png
// then AssetDir("data") would return []string{"foo.txt", "img"}
// AssetDir("data/img") would return []string{"a.png", "b.png"}
// AssetDir("foo.txt") and AssetDir("notexist") would return an error
// AssetDir("") will return []string{"data"}.
func AssetDir(name string) ([]string, error) {
	node := _bintree
	if len(name) != 0 {
		cannonicalName := strings.Replace(name, "\\", "/", -1)
		pathList := strings.Split(cannonicalName, "/")
		for _, p := range pathList {
			node = node.Children[p]
			if node == nil {
				return nil, fmt.Errorf("Asset %s not found", name)
			}
		}
	}
	if node.Func != nil {
		return nil, fmt.Errorf("Asset %s not found", name)
	}
	rv := make([]string, 0, len(node.Children))
	for childName := range node.Children {
		rv = append(rv, childName)
	}
	return rv, nil
}

type bintree struct {
	Func     func() (*asset, error)
	Children map[string]*bintree
}

var _bintree = &bintree{nil, map[string]*bintree{
	"stub": &bintree{nil, map[string]*bintree{
		"stub.json": &bintree{stubStubJson, map[string]*bintree{}},
	}},
}}

// RestoreAsset restores an asset under the given directory
func RestoreAsset(dir, name string) error {
	data, err := Asset(name)
	if err != nil {
		return err
	}
	info, err := AssetInfo(name)
	if err != nil {
		return err
	}
	err = os.MkdirAll(_filePath(dir, filepath.Dir(name)), os.FileMode(0755))
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(_filePath(dir, name), data, info.Mode())
	if err != nil {
		return err
	}
	err = os.Chtimes(_filePath(dir, name), info.ModTime(), info.ModTime())
	if err != nil {
		return err
	}
	return nil
}

// RestoreAssets restores an asset under the given directory recursively
func RestoreAssets(dir, name string) error {
	children, err := AssetDir(name)
	// File
	if err != nil {
		return RestoreAsset(dir, name)
	}
	// Dir
	for _, child := range children {
		err = RestoreAssets(dir, filepath.Join(name, child))
		if err != nil {
			return err
		}
	}
	return nil
}

func _filePath(dir, name string) string {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	return filepath.Join(append([]string{dir}, strings.Split(cannonicalName, "/")...)...)
}
