package parser

import (
	"strings"
)

const ParserNameOs = "os"
const FixtureFileOs = "oss.yml"

type OsReg struct {
	Regular `yaml:",inline" json:",inline"`
	Name    string `yaml:"name" json:"name"`
	Version string `yaml:"version" json:"version"`
}

// Known operating systems mapped to their internal short codes
var OperatingSystems = map[string]string{
	`AIX`: `AIX`,
	`AND`: `Android`,
	`ADR`: `Android TV`,
	`AMZ`: `Amazon Linux`,
	`AMG`: `AmigaOS`,
	`ATV`: `tvOS`,
	`ARL`: `Arch Linux`,
	`BTR`: `BackTrack`,
	`SBA`: `Bada`,
	`BEO`: `BeOS`,
	`BLB`: `BlackBerry OS`,
	`QNX`: `BlackBerry Tablet OS`,
	`BOS`: `Bliss OS`,
	`BMP`: `Brew`,
	`CAI`: `Caixa Mágica`,
	`CES`: `CentOS`,
	`CST`: `CentOS Stream`,
	`CLR`: `ClearOS Mobile`,
	`COS`: `Chrome OS`,
	`CRS`: `Chromium OS`,
	`CHN`: `China OS`,
	`CYN`: `CyanogenMod`,
	`DEB`: `Debian`,
	`DEE`: `Deepin`,
	`DFB`: `DragonFly`,
	`DVK`: `DVKBuntu`,
	`FED`: `Fedora`,
	`FEN`: `Fenix`,
	`FOS`: `Firefox OS`,
	`FIR`: `Fire OS`,
	`FOR`: `Foresight Linux`,
	`FRE`: `Freebox`,
	`BSD`: `FreeBSD`,
	`FYD`: `FydeOS`,
	`FUC`: `Fuchsia`,
	`GNT`: `Gentoo`,
	`GRI`: `GridOS`,
	`GTV`: `Google TV`,
	`HPX`: `HP-UX`,
	`HAI`: `Haiku OS`,
	`IPA`: `iPadOS`,
	`HAR`: `HarmonyOS`,
	`HAS`: `HasCodingOS`,
	`IRI`: `IRIX`,
	`INF`: `Inferno`,
	`JME`: `Java ME`,
	`KOS`: `KaiOS`,
	`KAN`: `Kanotix`,
	`KNO`: `Knoppix`,
	`KTV`: `KreaTV`,
	`KBT`: `Kubuntu`,
	`LIN`: `GNU/Linux`,
	`LND`: `LindowsOS`,
	`LNS`: `Linspire`,
	`LEN`: `Lineage OS`,
	`LBT`: `Lubuntu`,
	`LOS`: `Lumin OS`,
	`VLN`: `VectorLinux`,
	`MAC`: `Mac`,
	`MAE`: `Maemo`,
	`MAG`: `Mageia`,
	`MDR`: `Mandriva`,
	`SMG`: `MeeGo`,
	`MCD`: `MocorDroid`,
	`MON`: `moonOS`,
	`MIN`: `Mint`,
	`MLD`: `MildWild`,
	`MOR`: `MorphOS`,
	`NBS`: `NetBSD`,
	`MTK`: `MTK / Nucleus`,
	`MRE`: `MRE`,
	`WII`: `Nintendo`,
	`NDS`: `Nintendo Mobile`,
	`OS2`: `OS/2`,
	`T64`: `OSF1`,
	`OBS`: `OpenBSD`,
	`OWR`: `OpenWrt`,
	`OTV`: `Opera TV`,
	`ORD`: `Ordissimo`,
	`PAR`: `Pardus`,
	`PCL`: `PCLinuxOS`,
	`PLA`: `Plasma Mobile`,
	`PSP`: `PlayStation Portable`,
	`PS3`: `PlayStation`,
	`PUR`: `PureOS`,
	`RHT`: `Red Hat`,
	`REV`: `Revenge OS`,
	`ROS`: `RISC OS`,
	`ROK`: `Roku OS`,
	`RSO`: `Rosa`,
	`REM`: `Remix OS`,
	`REX`: `REX`,
	`RZD`: `RazoDroiD`,
	`SAB`: `Sabayon`,
	`SSE`: `SUSE`,
	`SAF`: `Sailfish OS`,
	`SEE`: `SeewoOS`,
	`SLW`: `Slackware`,
	`SOS`: `Solaris`,
	`SYL`: `Syllable`,
	`SYM`: `Symbian`,
	`SYS`: `Symbian OS`,
	`S40`: `Symbian OS Series 40`,
	`S60`: `Symbian OS Series 60`,
	`SY3`: `Symbian^3`,
	`TEN`: `TencentOS`,
	`TDX`: `ThreadX`,
	`TIZ`: `Tizen`,
	`TOS`: `TmaxOS`,
	`UBT`: `Ubuntu`,
	`WAS`: `watchOS`,
	`WTV`: `WebTV`,
	`WHS`: `Whale OS`,
	`WIN`: `Windows`,
	`WCE`: `Windows CE`,
	`WIO`: `Windows IoT`,
	`WMO`: `Windows Mobile`,
	`WPH`: `Windows Phone`,
	`WRT`: `Windows RT`,
	`XBX`: `Xbox`,
	`XBT`: `Xubuntu`,
	`YNS`: `YunOS`,
	`ZEN`: `Zenwalk`,
	`IOS`: `iOS`,
	`POS`: `palmOS`,
	`WOS`: `webOS`,
}

// Operating system families mapped to the short codes of the associated operating systems
var OsFamilies = map[string][]string{
	`Android`: []string{`AND`, `CYN`, `FIR`, `REM`, `RZD`, `MLD`, `MCD`, `YNS`, `GRI`, `HAR`,
		`ADR`, `CLR`, `BOS`, `REV`, `LEN`},
	`AmigaOS`:        {`AMG`, `MOR`},
	`Apple TV`:       {`ATV`},
	`BlackBerry`:     {`BLB`, `QNX`},
	`Brew`:           {`BMP`},
	`BeOS`:           {`BEO`, `HAI`},
	`Chrome OS`:      {`COS`, `CRS`, `FYD`, `SEE`},
	`Firefox OS`:     {`FOS`, `KOS`},
	`Gaming Console`: {`WII`, `PS3`},
	`Google TV`:      {`GTV`},
	`IBM`:            {`OS2`},
	`iOS`:            {`IOS`, `ATV`, `WAS`, `IPA`},
	`RISC OS`:        {`ROS`},
	`GNU/Linux`: {`LIN`, `ARL`, `DEB`, `KNO`, `MIN`, `UBT`, `KBT`, `XBT`, `LBT`, `FED`,
		`RHT`, `VLN`, `MDR`, `GNT`, `SAB`, `SLW`, `SSE`, `CES`, `BTR`, `SAF`,
		`ORD`, `TOS`, `RSO`, `DEE`, `FRE`, `MAG`, `FEN`, `CAI`, `PCL`, `HAS`,
		`LOS`, `DVK`, `ROK`, `OWR`, `OTV`, `KTV`, `PUR`, `PLA`, `FUC`, `PAR`,
		`FOR`, `MON`, `KAN`, `ZEN`, `LND`, `LNS`, `CHN`, `AMZ`, `TEN`, `CST`},
	`Mac`:                   {`MAC`},
	`Mobile Gaming Console`: {`PSP`, `NDS`, `XBX`},
	`Real-time OS`:          {`MTK`, `TDX`, `MRE`, `JME`, `REX`},
	`Other Mobile`:          {`WOS`, `POS`, `SBA`, `TIZ`, `SMG`, `MAE`},
	`Symbian`:               {`SYM`, `SYS`, `SY3`, `S60`, `S40`},
	`Unix`:                  {`SOS`, `AIX`, `HPX`, `BSD`, `NBS`, `OBS`, `DFB`, `SYL`, `IRI`, `T64`, `INF`},
	`WebTV`:                 {`WTV`},
	`Windows`:               {`WIN`},
	`Windows Mobile`:        {`WPH`, `WMO`, `WCE`, `WRT`, `WIO`},
	`Other Smart TV`:        {`WHS`},
}

var DesktopOS = map[string]bool{
	`AmigaOS`:     true,
	`IBM`:         true,
	`GNU/Linux`:   true,
	`Mac`:         true,
	`Unix`:        true,
	`Windows`:     true,
	`BeOS`:        true,
	`Chrome OS`:   true,
	`Chromium OS`: true,
}

const (
	PlatformTypeARM  = "ARM"
	PlatformTypeX64  = "x64"
	PlatformTypeX86  = "x86"
	PlatformTypeNONE = ""
)

type PlatformReg struct {
	Name string
	Regular
}

type OsMatchResult struct {
	Name      string `yaml:"name" json:"name"`
	ShortName string `yaml:"short_name" json:"short_name"`
	Version   string `yaml:"version" json:"version"`
	Platform  string `yaml:"platform" json:"platform"`
}

type OsParser interface {
	PreMatch(string) bool
	Parse(string) *OsMatchResult
}

// Parses the useragent for operating system information
type Oss struct {
	Regexes      []*OsReg
	platforms    []*PlatformReg
	overAllMatch Regular
}

func NewOss(file string) (*Oss, error) {
	var v []*OsReg
	err := ReadYamlFile(file, &v)
	if err != nil {
		return nil, err
	}
	ps := []*PlatformReg{
		{Name: PlatformTypeARM, Regular: Regular{Regex: "arm"}},
		{Name: PlatformTypeX64, Regular: Regular{Regex: "WOW64|x64|win64|amd64|x86_64"}},
		{Name: PlatformTypeX86, Regular: Regular{Regex: "i[0-9]86|i86pc"}},
	}
	for _, pp := range ps {
		pp.Compile()
	}
	return &Oss{
		Regexes:   v,
		platforms: ps,
	}, nil
}

func (o *Oss) ParsePlatform(ua string) string {
	for _, p := range o.platforms {
		if p.IsMatchUserAgent(ua) {
			return p.Name
		}
	}
	return PlatformTypeNONE
}

func (o *Oss) PreMatch(ua string) bool {
	if o.overAllMatch.Regexp == nil {
		count := len(o.Regexes)
		if count == 0 {
			return false
		}
		sb := strings.Builder{}
		sb.WriteString(o.Regexes[count-1].Regex)
		for i := count - 2; i >= 0; i-- {
			sb.WriteString("|")
			sb.WriteString(o.Regexes[i].Regex)
		}
		o.overAllMatch.Regex = sb.String()
		o.overAllMatch.Compile()
	}
	r := o.overAllMatch.IsMatchUserAgent(ua)
	return r
}

func (o *Oss) Parse(ua string) *OsMatchResult {
	var matches []string
	var osRegex *OsReg
	for _, osRegex = range o.Regexes {
		matches = osRegex.MatchUserAgent(ua)
		if len(matches) > 0 {
			break
		}
	}

	if len(matches) == 0 || osRegex == nil {
		return nil
	}

	name := BuildByMatch(osRegex.Name, matches)
	short := UnknownShort

	for osShort, osName := range OperatingSystems {
		if StringEqualIgnoreCase(name, osName) {
			name = osName
			short = osShort
			break
		}
	}

	result := &OsMatchResult{
		Name:      name,
		ShortName: short,
		Version:   BuildVersion(osRegex.Version, matches),
		Platform:  o.ParsePlatform(ua),
	}
	return result
}

func GetOsFamily(osLabel string) string {
	for k, vs := range OsFamilies {
		for _, v := range vs {
			if v == osLabel {
				return k
			}
		}
	}
	return ""
}

func GetOsNameFromId(os, ver string) string {
	if osFullName, ok := OperatingSystems[os]; ok {
		return strings.TrimSpace(osFullName + " " + ver)
	}
	return ""
}

func IsDesktopOS(os string) bool {
	osFamily := GetOsFamily(os)
	return DesktopOS[osFamily]
}
