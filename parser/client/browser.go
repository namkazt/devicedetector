package client

import (
	"path/filepath"

	gover "github.com/mcuadros/go-version"

	. "github.com/namkazt/devicedetector/parser"
)

// Known browsers mapped to their internal short codes
var availableBrowsers = map[string]string{
	`V1`: `Via`,
	`1P`: `Pure Mini Browser`,
	`4P`: `Pure Lite Browser`,
	`1R`: `Raise Fast Browser`,
	`FQ`: `Fast Browser UC Lite`,
	`FJ`: `Fast Explorer`,
	`1L`: `Lightning Browser`,
	`1C`: `Cake Browser`,
	`1I`: `IE Browser Fast`,
	`1V`: `Vegas Browser`,
	`1O`: `OH Browser`,
	`3O`: `OH Private Browser`,
	`1X`: `XBrowser Mini`,
	`1S`: `Sharkee Browser`,
	`2L`: `Lark Browser`,
	`3P`: `Pluma`,
	`1A`: `Anka Browser`,
	`AZ`: `Azka Browser`,
	`1D`: `Dragon Browser`,
	`1E`: `Easy Browser`,
	`DW`: `Dark Web Browser`,
	`1B`: `115 Browser`,
	`2B`: `2345 Browser`,
	`36`: `360 Phone Browser`,
	`3B`: `360 Browser`,
	`7B`: `7654 Browser`,
	`AA`: `Avant Browser`,
	`AB`: `ABrowse`,
	`BW`: `AdBlock Browser`,
	`AF`: `ANT Fresco`,
	`AG`: `ANTGalio`,
	`AL`: `Aloha Browser`,
	`AH`: `Aloha Browser Lite`,
	`AM`: `Amaya`,
	`A3`: `Amaze Browser`,
	`AO`: `Amigo`,
	`AN`: `Android Browser`,
	`AE`: `AOL Desktop`,
	`AD`: `AOL Shield`,
	`A4`: `AOL Shield Pro`,
	`AP`: `APUS Browser`,
	`AR`: `Arora`,
	`AX`: `Arctic Fox`,
	`AV`: `Amiga Voyager`,
	`AW`: `Amiga Aweb`,
	`AI`: `Arvin`,
	`AK`: `Ask.com`,
	`AU`: `Asus Browser`,
	`A0`: `Atom`,
	`AT`: `Atomic Web Browser`,
	`A2`: `Atlas`,
	`AS`: `Avast Secure Browser`,
	`VG`: `AVG Secure Browser`,
	`AC`: `Avira Scout`,
	`A1`: `AwoX`,
	`BA`: `Beaker Browser`,
	`BM`: `Beamrise`,
	`BB`: `BlackBerry Browser`,
	`H1`: `BrowseHere`,
	`BD`: `Baidu Browser`,
	`BS`: `Baidu Spark`,
	`BI`: `Basilisk`,
	`BV`: `Belva Browser`,
	`BE`: `Beonex`,
	`B2`: `Berry Browser`,
	`BT`: `Bitchute Browser`,
	`BH`: `BlackHawk`,
	`BJ`: `Bunjalloo`,
	`BL`: `B-Line`,
	`BU`: `Blue Browser`,
	`BO`: `Bonsai`,
	`BN`: `Borealis Navigator`,
	`BR`: `Brave`,
	`BK`: `BriskBard`,
	`BX`: `BrowseX`,
	`BZ`: `Browzar`,
	`BY`: `Biyubi`,
	`BF`: `Byffox`,
	`CA`: `Camino`,
	`CL`: `CCleaner`,
	`C8`: `CG Browser`,
	`CJ`: `ChanjetCloud`,
	`C6`: `Chedot`,
	`C0`: `Centaury`,
	`CC`: `Coc Coc`,
	`C4`: `CoolBrowser`,
	`C2`: `Colibri`,
	`CD`: `Comodo Dragon`,
	`C1`: `Coast`,
	`CX`: `Charon`,
	`CE`: `CM Browser`,
	`C7`: `CM Mini`,
	`CF`: `Chrome Frame`,
	`HC`: `Headless Chrome`,
	`CH`: `Chrome`,
	`CI`: `Chrome Mobile iOS`,
	`CK`: `Conkeror`,
	`CM`: `Chrome Mobile`,
	`3C`: `Chowbo`,
	`CN`: `CoolNovo`,
	`CO`: `CometBird`,
	`2C`: `Comfort Browser`,
	`CB`: `COS Browser`,
	`CW`: `Cornowser`,
	`C3`: `Chim Lac`,
	`CP`: `ChromePlus`,
	`CR`: `Chromium`,
	`C5`: `Chromium GOST`,
	`CY`: `Cyberfox`,
	`CS`: `Cheshire`,
	`CT`: `Crusta`,
	`CG`: `Craving Explorer`,
	`CZ`: `Crazy Browser`,
	`CU`: `Cunaguaro`,
	`CV`: `Chrome Webview`,
	`YC`: `CyBrowser`,
	`DB`: `dbrowser`,
	`PD`: `Peeps dBrowser`,
	`D1`: `Debuggable Browser`,
	`DC`: `Decentr`,
	`DE`: `Deepnet Explorer`,
	`DG`: `deg-degan`,
	`DA`: `Deledao`,
	`DT`: `Delta Browser`,
	`D0`: `Desi Browser`,
	`DS`: `DeskBrowse`,
	`DF`: `Dolphin`,
	`DZ`: `Dolphin Zero`,
	`DO`: `Dorado`,
	`DR`: `Dot Browser`,
	`DL`: `Dooble`,
	`DI`: `Dillo`,
	`DU`: `DUC Browser`,
	`DD`: `DuckDuckGo Privacy Browser`,
	`EC`: `Ecosia`,
	`EW`: `Edge WebView`,
	`EI`: `Epic`,
	`EL`: `Elinks`,
	`EN`: `EinkBro`,
	`EB`: `Element Browser`,
	`EE`: `Elements Browser`,
	`EZ`: `eZ Browser`,
	`EU`: `EUI Browser`,
	`EP`: `GNOME Web`,
	`ES`: `Espial TV Browser`,
	`FA`: `Falkon`,
	`FX`: `Faux Browser`,
	`F1`: `Firefox Mobile iOS`,
	`FB`: `Firebird`,
	`FD`: `Fluid`,
	`FE`: `Fennec`,
	`FF`: `Firefox`,
	`FK`: `Firefox Focus`,
	`FY`: `Firefox Reality`,
	`FR`: `Firefox Rocket`,
	`1F`: `Firefox Klar`,
	`F0`: `Float Browser`,
	`FL`: `Flock`,
	`FP`: `Floorp`,
	`FO`: `Flow`,
	`F2`: `Flow Browser`,
	`FM`: `Firefox Mobile`,
	`FW`: `Fireweb`,
	`FN`: `Fireweb Navigator`,
	`FH`: `Flash Browser`,
	`FS`: `Flast`,
	`FU`: `FreeU`,
	`F3`: `Frost+`,
	`FI`: `Fulldive`,
	`GA`: `Galeon`,
	`G8`: `Gener8`,
	`GH`: `Ghostery Privacy Browser`,
	`GI`: `GinxDroid Browser`,
	`GB`: `Glass Browser`,
	`GE`: `Google Earth`,
	`GP`: `Google Earth Pro`,
	`GO`: `GOG Galaxy`,
	`GR`: `GoBrowser`,
	`HB`: `Harman Browser`,
	`HS`: `HasBrowser`,
	`HA`: `Hawk Turbo Browser`,
	`HQ`: `Hawk Quick Browser`,
	`HE`: `Helio`,
	`HI`: `Hi Browser`,
	`HO`: `hola! Browser`,
	`HJ`: `HotJava`,
	`HU`: `Huawei Browser Mobile`,
	`HP`: `Huawei Browser`,
	`H3`: `HUB Browser`,
	`IO`: `iBrowser`,
	`IS`: `iBrowser Mini`,
	`IB`: `IBrowse`,
	`I6`: `iDesktop PC Browser`,
	`IC`: `iCab`,
	`I2`: `iCab Mobile`,
	`I1`: `Iridium`,
	`I3`: `Iron Mobile`,
	`I4`: `IceCat`,
	`ID`: `IceDragon`,
	`IV`: `Isivioo`,
	`IW`: `Iceweasel`,
	`IE`: `Internet Explorer`,
	`I5`: `Indian UC Mini Browser`,
	`IM`: `IE Mobile`,
	`IR`: `Iron`,
	`JB`: `Japan Browser`,
	`JS`: `Jasmine`,
	`JA`: `JavaFX`,
	`JL`: `Jelly`,
	`JI`: `Jig Browser`,
	`JP`: `Jig Browser Plus`,
	`JO`: `Jio Browser`,
	`J1`: `JioPages`,
	`KB`: `K.Browser`,
	`KS`: `Kids Safe Browser`,
	`KI`: `Kindle Browser`,
	`KM`: `K-meleon`,
	`KO`: `Konqueror`,
	`KP`: `Kapiko`,
	`KN`: `Kinza`,
	`KW`: `Kiwi`,
	`KD`: `Kode Browser`,
	`KT`: `KUTO Mini Browser`,
	`KY`: `Kylo`,
	`KZ`: `Kazehakase`,
	`LB`: `Cheetah Browser`,
	`LA`: `Lagatos Browser`,
	`LR`: `Lexi Browser`,
	`LV`: `Lenovo Browser`,
	`LF`: `LieBaoFast`,
	`LG`: `LG Browser`,
	`LH`: `Light`,
	`L1`: `Lilo`,
	`LI`: `Links`,
	`IF`: `Lolifox`,
	`LO`: `Lovense Browser`,
	`LT`: `LT Browser`,
	`LU`: `LuaKit`,
	`LL`: `Lulumi`,
	`LS`: `Lunascape`,
	`LN`: `Lunascape Lite`,
	`LX`: `Lynx`,
	`MD`: `Mandarin`,
	`M1`: `mCent`,
	`MB`: `MicroB`,
	`MC`: `NCSA Mosaic`,
	`MZ`: `Meizu Browser`,
	`ME`: `Mercury`,
	`M2`: `Me Browser`,
	`MF`: `Mobile Safari`,
	`MI`: `Midori`,
	`M3`: `Midori Lite`,
	`MO`: `Mobicip`,
	`MU`: `MIUI Browser`,
	`MS`: `Mobile Silk`,
	`MN`: `Minimo`,
	`MT`: `Mint Browser`,
	`MX`: `Maxthon`,
	`M4`: `MaxTube Browser`,
	`MA`: `Maelstrom`,
	`MM`: `Mmx Browser`,
	`NM`: `MxNitro`,
	`MY`: `Mypal`,
	`MR`: `Monument Browser`,
	`MW`: `MAUI WAP Browser`,
	`NA`: `Navegador`,
	`NW`: `Navigateur Web`,
	`NK`: `Naked Browser`,
	`NR`: `NFS Browser`,
	`NB`: `Nokia Browser`,
	`NO`: `Nokia OSS Browser`,
	`NV`: `Nokia Ovi Browser`,
	`NX`: `Nox Browser`,
	`NE`: `NetSurf`,
	`NF`: `NetFront`,
	`NL`: `NetFront Life`,
	`NP`: `NetPositive`,
	`NS`: `Netscape`,
	`NT`: `NTENT Browser`,
	`OC`: `Oculus Browser`,
	`O1`: `Opera Mini iOS`,
	`OB`: `Obigo`,
	`O2`: `Odin`,
	`2O`: `Odin Browser`,
	`H2`: `OceanHero`,
	`OD`: `Odyssey Web Browser`,
	`OF`: `Off By One`,
	`HH`: `OhHai Browser`,
	`OE`: `ONE Browser`,
	`Y1`: `Opera Crypto`,
	`OX`: `Opera GX`,
	`OG`: `Opera Neon`,
	`OH`: `Opera Devices`,
	`OI`: `Opera Mini`,
	`OM`: `Opera Mobile`,
	`OP`: `Opera`,
	`ON`: `Opera Next`,
	`OO`: `Opera Touch`,
	`OA`: `Orca`,
	`OS`: `Ordissimo`,
	`OR`: `Oregano`,
	`O0`: `Origin In-Game Overlay`,
	`OY`: `Origyn Web Browser`,
	`OV`: `Openwave Mobile Browser`,
	`O3`: `OpenFin`,
	`O4`: `Open Browser`,
	`4U`: `Open Browser 4U`,
	`OW`: `OmniWeb`,
	`OT`: `Otter Browser`,
	`PL`: `Palm Blazer`,
	`PM`: `Pale Moon`,
	`PY`: `Polypane`,
	`PP`: `Oppo Browser`,
	`PR`: `Palm Pre`,
	`PU`: `Puffin`,
	`2P`: `Puffin Web Browser`,
	`PW`: `Palm WebPro`,
	`PA`: `Palmscape`,
	`PE`: `Perfect Browser`,
	`P1`: `Phantom.me`,
	`PH`: `Phantom Browser`,
	`PX`: `Phoenix`,
	`PB`: `Phoenix Browser`,
	`PF`: `PlayFree Browser`,
	`PK`: `PocketBook Browser`,
	`PO`: `Polaris`,
	`PT`: `Polarity`,
	`LY`: `PolyBrowser`,
	`PI`: `PrivacyWall`,
	`P0`: `PronHub Browser`,
	`PC`: `PSI Secure Browser`,
	`RW`: `Reqwireless WebViewer`,
	`PS`: `Microsoft Edge`,
	`QA`: `Qazweb`,
	`Q2`: `QQ Browser Lite`,
	`Q1`: `QQ Browser Mini`,
	`QQ`: `QQ Browser`,
	`QS`: `Quick Browser`,
	`QT`: `Qutebrowser`,
	`QU`: `Quark`,
	`QZ`: `QupZilla`,
	`QM`: `Qwant Mobile`,
	`QW`: `QtWebEngine`,
	`RE`: `Realme Browser`,
	`RK`: `Rekonq`,
	`RM`: `RockMelt`,
	`SB`: `Samsung Browser`,
	`SA`: `Sailfish Browser`,
	`S8`: `Seewo Browser`,
	`SC`: `SEMC-Browser`,
	`SE`: `Sogou Explorer`,
	`SO`: `Sogou Mobile Browser`,
	`RF`: `SOTI Surf`,
	`2S`: `Soul Browser`,
	`SF`: `Safari`,
	`PV`: `Safari Technology Preview`,
	`S5`: `Safe Exam Browser`,
	`SW`: `SalamWeb`,
	`VN`: `Savannah Browser`,
	`SD`: `SavySoda`,
	`S9`: `Secure Browser`,
	`SV`: `SFive`,
	`SH`: `Shiira`,
	`K1`: `Sidekick`,
	`S1`: `SimpleBrowser`,
	`3S`: `SilverMob US`,
	`SY`: `Sizzy`,
	`SK`: `Skyfire`,
	`SS`: `Seraphic Sraf`,
	`KK`: `SiteKiosk`,
	`SL`: `Sleipnir`,
	`S6`: `Slimjet`,
	`S7`: `SP Browser`,
	`8S`: `Secure Private Browser`,
	`T1`: `Stampy Browser`,
	`7S`: `7Star`,
	`SQ`: `Smart Browser`,
	`6S`: `Smart Search & Web Browser`,
	`LE`: `Smart Lenovo Browser`,
	`OZ`: `Smooz`,
	`SN`: `Snowshoe`,
	`B1`: `Spectre Browser`,
	`S2`: `Splash`,
	`SI`: `Sputnik Browser`,
	`SR`: `Sunrise`,
	`SP`: `SuperBird`,
	`SU`: `Super Fast Browser`,
	`5S`: `SuperFast Browser`,
	`HR`: `Sushi Browser`,
	`S3`: `surf`,
	`4S`: `Surf Browser`,
	`SG`: `Stargon`,
	`S0`: `START Internet Browser`,
	`S4`: `Steam In-Game Overlay`,
	`ST`: `Streamy`,
	`SX`: `Swiftfox`,
	`SZ`: `Seznam Browser`,
	`TP`: `T+Browser`,
	`TR`: `T-Browser`,
	`TO`: `t-online.de Browser`,
	`TA`: `Tao Browser`,
	`TF`: `TenFourFox`,
	`TB`: `Tenta Browser`,
	`TE`: `Tesla Browser`,
	`TZ`: `Tizen Browser`,
	`TU`: `Tungsten`,
	`TG`: `ToGate`,
	`TS`: `TweakStyle`,
	`TV`: `TV Bro`,
	`UB`: `UBrowser`,
	`UC`: `UC Browser`,
	`UH`: `UC Browser HD`,
	`UM`: `UC Browser Mini`,
	`UT`: `UC Browser Turbo`,
	`UI`: `Ui Browser Mini`,
	`UR`: `UR Browser`,
	`UZ`: `Uzbl`,
	`UE`: `Ume Browser`,
	`V0`: `vBrowser`,
	`VE`: `Venus Browser`,
	`N0`: `Nova Video Downloader Pro`,
	`VS`: `Viasat Browser`,
	`VI`: `Vivaldi`,
	`VV`: `vivo Browser`,
	`VB`: `Vision Mobile Browser`,
	`VM`: `VMware AirWatch`,
	`WI`: `Wear Internet Browser`,
	`WP`: `Web Explorer`,
	`WE`: `WebPositive`,
	`WF`: `Waterfox`,
	`WB`: `Wave Browser`,
	`WH`: `Whale Browser`,
	`WO`: `wOSBrowser`,
	`WT`: `WeTab Browser`,
	`YJ`: `Yahoo! Japan Browser`,
	`YA`: `Yandex Browser`,
	`YL`: `Yandex Browser Lite`,
	`YN`: `Yaani Browser`,
	`YB`: `Yolo Browser`,
	`YO`: `YouCare`,
	`XS`: `xStand`,
	`XI`: `Xiino`,
	`XV`: `Xvast`,
	`ZE`: `Zetakey`,
	`ZV`: `Zvu`,

	// detected browsers in older versions
	// `IA` : `Iceape`,  => pim
	// `SM` : `SeaMonkey`,  => pim
}

// Browser families mapped to the short codes of the associated browsers
var browserFamilies = map[string][]string{
	`Android Browser`:    {`AN`, `MU`},
	`BlackBerry Browser`: {`BB`},
	`Baidu`:              {`BD`, `BS`},
	`Amiga`:              {`AV`, `AW`},
	`Chrome`: {`1B`, `2B`, `7S`, `A0`, `AC`, `A4`, `AE`, `AH`, `AI`,
		`AO`, `AS`, `BA`, `BM`, `BR`, `C2`, `C3`, `C5`, `C4`,
		`C6`, `CC`, `CD`, `CE`, `CF`, `CG`, `CH`, `CI`, `CL`,
		`CM`, `CN`, `CP`, `CR`, `CV`, `CW`, `DA`, `DD`, `DG`,
		`DR`, `EC`, `EE`, `EU`, `EW`, `FA`, `FS`, `GB`, `GI`,
		`H2`, `HA`, `HE`, `HH`, `HS`, `I3`, `IR`, `JB`, `KN`,
		`KW`, `LF`, `LL`, `LO`, `M1`, `MA`, `MD`, `MR`, `MS`,
		`MT`, `MZ`, `NM`, `NR`, `O0`, `O2`, `O3`, `OC`, `PB`,
		`PT`, `QU`, `QW`, `RM`, `S4`, `S6`, `S8`, `S9`, `SB`,
		`SG`, `SS`, `SU`, `SV`, `SW`, `SY`, `SZ`, `T1`, `TA`,
		`TB`, `TG`, `TR`, `TS`, `TU`, `TV`, `UB`, `UR`, `VE`,
		`VG`, `VI`, `VM`, `WP`, `WH`, `XV`, `YJ`, `YN`, `FH`,
		`B1`, `BO`, `HB`, `PC`, `LA`, `LT`, `PD`, `HR`, `HU`,
		`HP`, `IO`, `TP`, `CJ`, `HQ`, `HI`, `NA`, `BW`, `YO`,
		`DC`, `G8`, `DT`, `AP`, `AK`, `UI`, `SD`, `VN`, `4S`,
		`2S`, `RF`, `LR`, `SQ`, `BV`, `L1`, `F0`, `KS`, `V0`,
		`C8`, `AZ`, `MM`, `BT`, `N0`, `P0`, `F3`, `VS`, `DU`,
		`D0`, `P1`, `O4`, `8S`, `H3`, `TE`, `WB`, `K1`},
	`Firefox`: {`AX`, `BI`, `BF`, `BH`, `BN`, `C0`, `CU`, `EI`, `F1`,
		`FB`, `FE`, `FF`, `FM`, `FR`, `FY`, `GZ`, `I4`, `IF`,
		`IW`, `LH`, `LY`, `MB`, `MN`, `MO`, `MY`, `OA`, `OS`,
		`PI`, `PX`, `QA`, `QM`, `S5`, `SX`, `TF`, `TO`, `WF`,
		`ZV`, `FP`, `AD`},
	`Internet Explorer`: {`BZ`, `CZ`, `IE`, `IM`, `PS`},
	`Konqueror`:         {`KO`},
	`NetFront`:          {`NF`},
	`NetSurf`:           {`NE`},
	`Nokia Browser`:     {`NB`, `NO`, `NV`, `DO`},
	`Opera`:             {`O1`, `OG`, `OH`, `OI`, `OM`, `ON`, `OO`, `OP`, `OX`, `Y1`},
	`Safari`:            {`MF`, `S7`, `SF`, `SO`, `PV`},
	`Sailfish Browser`:  {`SA`},
}

// Browsers that are available for mobile devices only
var mobileOnlyBrowsers = []string{
	`36`, `AH`, `AI`, `BL`, `C1`, `C4`, `CB`, `CW`, `DB`,
	`DD`, `DT`, `EU`, `EZ`, `FK`, `FM`, `FR`, `FX`, `GH`,
	`GI`, `GR`, `HA`, `HU`, `IV`, `JB`, `KD`, `M1`, `MF`,
	`MN`, `MZ`, `NX`, `OC`, `OI`, `OM`, `OZ`, `PU`, `PI`,
	`PE`, `QU`, `RE`, `S0`, `S7`, `SA`, `SB`, `SG`, `SK`,
	`ST`, `SU`, `T1`, `UH`, `UM`, `UT`, `VE`, `VV`, `WI`,
	`WP`, `YN`, `IO`, `IS`, `HQ`, `RW`, `HI`, `NA`, `BW`,
	`YO`, `PK`, `MR`, `AP`, `AK`, `UI`, `SD`, `VN`, `4S`,
	`RF`, `LR`, `SQ`, `BV`, `L1`, `F0`, `KS`, `V0`, `C8`,
	`AZ`, `MM`, `BT`, `N0`, `P0`, `F3`, `DU`, `D0`, `P1`,
	`O4`,
}

func GetBrowserFamily(browserLabel string) (string, bool) {
	for k, vs := range browserFamilies {
		for _, v := range vs {
			if v == browserLabel {
				return k, true
			}
		}
	}
	return "", false
}

// Returns if the given browser is mobile only
func IsMobileOnlyBrowser(browser string) bool {
	if ArrayContainsString(mobileOnlyBrowsers, browser) {
		return true
	}
	if v, ok := availableBrowsers[browser]; ok {
		return ArrayContainsString(mobileOnlyBrowsers, v)
	}
	return false
}

type BrowserMatchResult = ClientMatchResult

type Engine struct {
	Default  string            `yaml:"default" json:"default"`
	Versions map[string]string `yaml:"versions" json:"versions"`
}

type BrowserItem struct {
	Regular `yaml:",inline" json:",inline"`
	Name    string  `yaml:"name" json:"name"`
	Version string  `yaml:"version" json:"version"`
	Engine  *Engine `yaml:"engine" json:"engine"`
}

//  Client parser for browser detection
type Browser struct {
	Regexes  []*BrowserItem
	engine   BrowserEngine
	verCache map[string]*Version
}

const ParserNameBrowser = `browser`
const FixtureFileBrowser = `browsers.yml`

func init() {
	RegClientParser(ParserNameBrowser,
		func(dir string) ClientParser {
			return NewBrowser(filepath.Join(dir, FixtureFileBrowser))
		})
}

func NewBrowser(fileName string) *Browser {
	b := &Browser{}
	b.engine.ParserName = ParserNameBrowserEngine
	if err := b.Load(fileName); err != nil {
		return nil
	}
	return b
}

func (b *Browser) Load(file string) error {
	b.verCache = make(map[string]*Version)
	var v []*BrowserItem
	err := ReadYamlFile(file, &v)
	if err != nil {
		return err
	}
	engineFile := file[0:len(file)-len(FixtureFileBrowser)] + FixtureFileBrowserEngine
	err = b.engine.Load(engineFile)
	if err != nil {
		return err
	}
	b.Regexes = v
	return nil
}

func (b *Browser) PreMatch(ua string) bool {
	return true
}

func (b *Browser) Parse(ua string) *BrowserMatchResult {
	for _, regex := range b.Regexes {
		matches := regex.MatchUserAgent(ua)
		if len(matches) > 0 {
			name := BuildByMatch(regex.Name, matches)
			for browserShort, browserName := range availableBrowsers {
				if StringEqualIgnoreCase(name, browserName) {
					version := BuildVersion(regex.Version, matches)
					engine := b.BuildEngine(regex.Engine, version, ua)
					engineVersion := b.BuildEngineVersion(engine, ua)
					return &BrowserMatchResult{
						Type:          ParserNameBrowser,
						Name:          browserName,
						ShortName:     browserShort,
						Version:       version,
						Engine:        engine,
						EngineVersion: engineVersion,
					}
				}
			}
		}
	}
	return nil
}

func (b *Browser) BuildEngine(engineData *Engine, browserVersion, ua string) string {
	engine := ""
	if engineData != nil {
		engine = engineData.Default
		for version, versionEngine := range engineData.Versions {
			if gover.CompareSimple(browserVersion, version) >= 0 {
				engine = versionEngine
			}
		}
	}
	if engine == "" {
		if engineResult := b.engine.Parse(ua); engineResult != nil {
			engine = engineResult.Name
		}
	}
	return engine
}

func (b *Browser) BuildEngineVersion(engine, ua string) string {
	if engine == "" {
		return ""
	}
	v, ok := b.verCache[engine]
	if !ok {
		v = &Version{Engine: engine}
		v.Compile()
		b.verCache[engine] = v
	}
	return v.Parse(ua)
}
