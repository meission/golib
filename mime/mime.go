/*------------------------
name        mime
describe    meme type
author      ailn(z.ailn@wmntec.com)
create      2016-05-08
version     1.0
------------------------*/
package mime

import (
	//golang offical package
	"bytes"
	"encoding/hex"
)

type Mime struct {
	Suffix string
	Type   string
}

//give the suffix of file to get mime
func Suffix(suffix string) Mime {
	for _, m := range Mimes {
		if suffix == m.Suffix {
			return m
		}
	}
	return UKNOWN
}

//give the first two Byte to get mime
func FileHeader(file []byte) Mime {
	for k, v := range Headers {
		header, err := hex.DecodeString(k)
		if err != nil {
			continue
		}
		if bytes.HasPrefix(file, header) {
			return v
		}
	}
	return UKNOWN
}

//file headers map to mime
var Headers = map[string]Mime{
	hex.EncodeToString([]byte{255, 216}): IMAGEJPG,
	hex.EncodeToString([]byte{71, 73}):   IMAGEGIF,
	hex.EncodeToString([]byte{66, 77}):   IMAGEBMP,
	hex.EncodeToString([]byte{137, 80}):  IMAGEPNG,
	hex.EncodeToString([]byte{208, 207}): APPDOC,
	hex.EncodeToString([]byte{80, 75}):   APPDOCX,
	hex.EncodeToString([]byte{208, 207}): APPXLS,
	hex.EncodeToString([]byte{80, 75}):   APPXLSX,
	hex.EncodeToString([]byte{239, 187}): APPJS,
	hex.EncodeToString([]byte{67, 87}):   APPSWF,
	hex.EncodeToString([]byte{70, 67}):   TEXTTXT,
	hex.EncodeToString([]byte{73, 68}):   AUDIOMP3,
	hex.EncodeToString([]byte{48, 38}):   AUDIOWMA,
	hex.EncodeToString([]byte{77, 84}):   AUDIOMID,
	hex.EncodeToString([]byte{82, 97}):   APPRAR,
	hex.EncodeToString([]byte{80, 75}):   APPZIP,
	hex.EncodeToString([]byte{60, 63}):   APPXML,
}

var Mimes = []Mime{
	Mime{".323", "text/h323"},
	Mime{".3gp", "video/3gpp"},
	Mime{".aab", "application/x-authoware-bin"},
	Mime{".aam", "application/x-authoware-map"},
	Mime{".aas", "application/x-authoware-seg"},
	Mime{".acx", "application/internet-property-stream"},
	Mime{".ai", "application/postscript"},
	Mime{".aif", "audio/x-aiff"},
	Mime{".aifc", "audio/x-aiff"},
	Mime{".aiff", "audio/x-aiff"},
	Mime{".als", "audio/X-Alpha5"},
	Mime{".amc", "application/x-mpeg"},
	Mime{".ani", "application/octet-stream"},
	Mime{".apk", "application/vnd.android.package-archive"},
	Mime{".asc", "text/plain"},
	Mime{".asd", "application/astound"},
	Mime{".asf", "video/x-ms-asf"},
	Mime{".asn", "application/astound"},
	Mime{".asp", "application/x-asap"},
	Mime{".asr", "video/x-ms-asf"},
	Mime{".asx", "video/x-ms-asf"},
	Mime{".au", "audio/basic"},
	Mime{".avb", "application/octet-stream"},
	Mime{".avi", "video/x-msvideo"},
	Mime{".awb", "audio/amr-wb"},
	Mime{".axs", "application/olescript"},
	Mime{".bas", "text/plain"},
	Mime{".bcpio", "application/x-bcpio"},
	Mime{".bin ", "application/octet-stream"},
	Mime{".bld", "application/bld"},
	Mime{".bld2", "application/bld2"},
	Mime{".bmp", "image/bmp"},
	Mime{".bpk", "application/octet-stream"},
	Mime{".bz2", "application/x-bzip2"},
	Mime{".c", "text/plain"},
	Mime{".cal", "image/x-cals"},
	Mime{".cat", "application/vnd.ms-pkiseccat"},
	Mime{".ccn", "application/x-cnc"},
	Mime{".cco", "application/x-cocoa"},
	Mime{".cdf", "application/x-cdf"},
	Mime{".cer", "application/x-x509-ca-cert"},
	Mime{".cgi", "magnus-internal/cgi"},
	Mime{".chat", "application/x-chat"},
	Mime{".class", "application/octet-stream"},
	Mime{".clp", "application/x-msclip"},
	Mime{".cmx", "image/x-cmx"},
	Mime{".co", "application/x-cult3d-object"},
	Mime{".cod", "image/cis-cod"},
	Mime{".conf", "text/plain"},
	Mime{".cpio", "application/x-cpio"},
	Mime{".cpp", "text/plain"},
	Mime{".cpt", "application/mac-compactpro"},
	Mime{".crd", "application/x-mscardfile"},
	Mime{".crl", "application/pkix-crl"},
	Mime{".crt", "application/x-x509-ca-cert"},
	Mime{".csh", "application/x-csh"},
	Mime{".csm", "chemical/x-csml"},
	Mime{".csml", "chemical/x-csml"},
	Mime{".css", "text/css"},
	Mime{".cur", "application/octet-stream"},
	Mime{".dcm", "x-lml/x-evm"},
	Mime{".dcr", "application/x-director"},
	Mime{".dcx", "image/x-dcx"},
	Mime{".der", "application/x-x509-ca-cert"},
	Mime{".dhtml", "text/html"},
	Mime{".dir", "application/x-director"},
	Mime{".dll", "application/x-msdownload"},
	Mime{".dmg", "application/octet-stream"},
	Mime{".dms", "application/octet-stream"},
	Mime{".doc", "application/msword"},
	Mime{".docx", "application/vnd.openxmlformats-officedocument.wordprocessingml.document"},
	Mime{".dot", "application/msword"},
	Mime{".dvi", "application/x-dvi"},
	Mime{".dwf", "drawing/x-dwf"},
	Mime{".dwg", "application/x-autocad"},
	Mime{".dxf", "application/x-autocad"},
	Mime{".dxr", "application/x-director"},
	Mime{".ebk", "application/x-expandedbook"},
	Mime{".emb", "chemical/x-embl-dl-nucleotide"},
	Mime{".embl", "chemical/x-embl-dl-nucleotide"},
	Mime{".eps", "application/postscript"},
	Mime{".epub", "application/epub+zip"},
	Mime{".eri", "image/x-eri"},
	Mime{".es", "audio/echospeech"},
	Mime{".esl", "audio/echospeech"},
	Mime{".etc", "application/x-earthtime"},
	Mime{".etx", "text/x-setext"},
	Mime{".evm", "x-lml/x-evm"},
	Mime{".evy", "application/envoy"},
	Mime{".exe", "application/octet-stream"},
	Mime{".fh4", "image/x-freehand"},
	Mime{".fh5", "image/x-freehand"},
	Mime{".fhc", "image/x-freehand"},
	Mime{".fif", "application/fractals"},
	Mime{".flr", "x-world/x-vrml"},
	Mime{".flv", "flv-application/octet-stream"},
	Mime{".fm", "application/x-maker"},
	Mime{".fpx", "image/x-fpx"},
	Mime{".fvi", "video/isivideo"},
	Mime{".gau", "chemical/x-gaussian-input"},
	Mime{".gca", "application/x-gca-compressed"},
	Mime{".gdb", "x-lml/x-gdb"},
	Mime{".gif", "image/gif"},
	Mime{".gps", "application/x-gps"},
	Mime{".gtar", "application/x-gtar"},
	Mime{".gz", "application/x-gzip"},
	Mime{".h", "text/plain"},
	Mime{".hdf", "application/x-hdf"},
	Mime{".hdm", "text/x-hdml"},
	Mime{".hdml", "text/x-hdml"},
	Mime{".hlp", "application/winhlp"},
	Mime{".hqx", "application/mac-binhex40"},
	Mime{".hta", "application/hta"},
	Mime{".htc", "text/x-component"},
	Mime{".htm", "text/html"},
	Mime{".html", "text/html"},
	Mime{".hts", "text/html"},
	Mime{".htt", "text/webviewhtml"},
	Mime{".ice", "x-conference/x-cooltalk"},
	Mime{".ico", "image/x-icon"},
	Mime{".ief", "image/ief"},
	Mime{".ifm", "image/gif"},
	Mime{".ifs", "image/ifs"},
	Mime{".iii", "application/x-iphone"},
	Mime{".imy", "audio/melody"},
	Mime{".ins", "application/x-internet-signup"},
	Mime{".ips", "application/x-ipscript"},
	Mime{".ipx", "application/x-ipix"},
	Mime{".isp", "application/x-internet-signup"},
	Mime{".it", "audio/x-mod"},
	Mime{".itz", "audio/x-mod"},
	Mime{".ivr", "i-world/i-vrml"},
	Mime{".j2k", "image/j2k"},
	Mime{".jad", "text/vnd.sun.j2me.app-descriptor"},
	Mime{".jam", "application/x-jam"},
	Mime{".jar", "application/java-archive"},
	Mime{".java", "text/plain"},
	Mime{".jfif", "image/pipeg"},
	Mime{".jnlp", "application/x-java-jnlp-file"},
	Mime{".jpe", "image/jpeg"},
	Mime{".jpeg", "image/jpeg"},
	Mime{".jpg", "image/jpeg"},
	Mime{".jpz", "image/jpeg"},
	Mime{".js", "application/x-javascript"},
	Mime{".jwc", "application/jwc"},
	Mime{".kjx", "application/x-kjx"},
	Mime{".lak", "x-lml/x-lak"},
	Mime{".latex", "application/x-latex"},
	Mime{".lcc", "application/fastman"},
	Mime{".lcl", "application/x-digitalloca"},
	Mime{".lcr", "application/x-digitalloca"},
	Mime{".lgh", "application/lgh"},
	Mime{".lha", "application/octet-stream"},
	Mime{".lml", "x-lml/x-lml"},
	Mime{".lmlpack", "x-lml/x-lmlpack"},
	Mime{".log", "text/plain"},
	Mime{".lsf", "video/x-la-asf"},
	Mime{".lsx", "video/x-la-asf"},
	Mime{".lzh", "application/octet-stream"},
	Mime{".m13", "application/x-msmediaview"},
	Mime{".m14", "application/x-msmediaview"},
	Mime{".m15", "audio/x-mod"},
	Mime{".m3u", "audio/x-mpegurl"},
	Mime{".m3url", "audio/x-mpegurl"},
	Mime{".m4a", "audio/mp4a-latm"},
	Mime{".m4b", "audio/mp4a-latm"},
	Mime{".m4p", "audio/mp4a-latm"},
	Mime{".m4u", "video/vnd.mpegurl"},
	Mime{".m4v", "video/x-m4v"},
	Mime{".ma1", "audio/ma1"},
	Mime{".ma2", "audio/ma2"},
	Mime{".ma3", "audio/ma3"},
	Mime{".ma5", "audio/ma5"},
	Mime{".man", "application/x-troff-man"},
	Mime{".map", "magnus-internal/imagemap"},
	Mime{".mbd", "application/mbedlet"},
	Mime{".mct", "application/x-mascot"},
	Mime{".mdb", "application/x-msaccess"},
	Mime{".mdz", "audio/x-mod"},
	Mime{".me", "application/x-troff-me"},
	Mime{".mel", "text/x-vmel"},
	Mime{".mht", "message/rfc822"},
	Mime{".mhtml", "message/rfc822"},
	Mime{".mi", "application/x-mif"},
	Mime{".mid", "audio/mid"},
	Mime{".midi", "audio/midi"},
	Mime{".mif", "application/x-mif"},
	Mime{".mil", "image/x-cals"},
	Mime{".mio", "audio/x-mio"},
	Mime{".mmf", "application/x-skt-lbs"},
	Mime{".mng", "video/x-mng"},
	Mime{".mny", "application/x-msmoney"},
	Mime{".moc", "application/x-mocha"},
	Mime{".mocha", "application/x-mocha"},
	Mime{".mod", "audio/x-mod"},
	Mime{".mof", "application/x-yumekara"},
	Mime{".mol", "chemical/x-mdl-molfile"},
	Mime{".mop", "chemical/x-mopac-input"},
	Mime{".mov", "video/quicktime"},
	Mime{".movie", "video/x-sgi-movie"},
	Mime{".mp2", "video/mpeg"},
	Mime{".mp3", "audio/mpeg"},
	Mime{".mp4", "video/mp4"},
	Mime{".mpa", "video/mpeg"},
	Mime{".mpc", "application/vnd.mpohun.certificate"},
	Mime{".mpe", "video/mpeg"},
	Mime{".mpeg", "video/mpeg"},
	Mime{".mpg", "video/mpeg"},
	Mime{".mpg4", "video/mp4"},
	Mime{".mpga", "audio/mpeg"},
	Mime{".mpn", "application/vnd.mophun.application"},
	Mime{".mpp", "application/vnd.ms-project"},
	Mime{".mps", "application/x-mapserver"},
	Mime{".mpv2", "video/mpeg"},
	Mime{".mrl", "text/x-mrml"},
	Mime{".mrm", "application/x-mrm"},
	Mime{".ms", "application/x-troff-ms"},
	Mime{".msg", "application/vnd.ms-outlook"},
	Mime{".mts", "application/metastream"},
	Mime{".mtx", "application/metastream"},
	Mime{".mtz", "application/metastream"},
	Mime{".mvb", "application/x-msmediaview"},
	Mime{".mzv", "application/metastream"},
	Mime{".nar", "application/zip"},
	Mime{".nbmp", "image/nbmp"},
	Mime{".nc", "application/x-netcdf"},
	Mime{".ndb", "x-lml/x-ndb"},
	Mime{".ndwn", "application/ndwn"},
	Mime{".nif", "application/x-nif"},
	Mime{".nmz", "application/x-scream"},
	Mime{".nokia-op-logo", "image/vnd.nok-oplogo-color"},
	Mime{".npx", "application/x-netfpx"},
	Mime{".nsnd", "audio/nsnd"},
	Mime{".nva", "application/x-neva1"},
	Mime{".nws", "message/rfc822"},
	Mime{".oda", "application/oda"},
	Mime{".ogg", "audio/ogg"},
	Mime{".oom", "application/x-AtlasMate-Plugin"},
	Mime{".p10", "application/pkcs10"},
	Mime{".p12", "application/x-pkcs12"},
	Mime{".p7b", "application/x-pkcs7-certificates"},
	Mime{".p7c", "application/x-pkcs7-mime"},
	Mime{".p7m", "application/x-pkcs7-mime"},
	Mime{".p7r", "application/x-pkcs7-certreqresp"},
	Mime{".p7s", "application/x-pkcs7-signature"},
	Mime{".pac", "audio/x-pac"},
	Mime{".pae", "audio/x-epac"},
	Mime{".pan", "application/x-pan"},
	Mime{".pbm", "image/x-portable-bitmap"},
	Mime{".pcx", "image/x-pcx"},
	Mime{".pda", "image/x-pda"},
	Mime{".pdb", "chemical/x-pdb"},
	Mime{".pdf", "application/pdf"},
	Mime{".pfr", "application/font-tdpfr"},
	Mime{".pfx", "application/x-pkcs12"},
	Mime{".pgm", "image/x-portable-graymap"},
	Mime{".pict", "image/x-pict"},
	Mime{".pko", "application/ynd.ms-pkipko"},
	Mime{".pm", "application/x-perl"},
	Mime{".pma", "application/x-perfmon"},
	Mime{".pmc", "application/x-perfmon"},
	Mime{".pmd", "application/x-pmd"},
	Mime{".pml", "application/x-perfmon"},
	Mime{".pmr", "application/x-perfmon"},
	Mime{".pmw", "application/x-perfmon"},
	Mime{".png", "image/png"},
	Mime{".pnm", "image/x-portable-anymap"},
	Mime{".pnz", "image/png"},
	Mime{".pot,", "application/vnd.ms-powerpoint"},
	Mime{".ppm", "image/x-portable-pixmap"},
	Mime{".pps", "application/vnd.ms-powerpoint"},
	Mime{".ppt", "application/vnd.ms-powerpoint"},
	Mime{".pptx", "application/vnd.openxmlformats-officedocument.presentationml.presentation"},
	Mime{".pqf", "application/x-cprplayer"},
	Mime{".pqi", "application/cprplayer"},
	Mime{".prc", "application/x-prc"},
	Mime{".prf", "application/pics-rules"},
	Mime{".prop", "text/plain"},
	Mime{".proxy", "application/x-ns-proxy-autoconfig"},
	Mime{".ps", "application/postscript"},
	Mime{".ptlk", "application/listenup"},
	Mime{".pub", "application/x-mspublisher"},
	Mime{".pvx", "video/x-pv-pvx"},
	Mime{".qcp", "audio/vnd.qcelp"},
	Mime{".qt", "video/quicktime"},
	Mime{".qti", "image/x-quicktime"},
	Mime{".qtif", "image/x-quicktime"},
	Mime{".r3t", "text/vnd.rn-realtext3d"},
	Mime{".ra", "audio/x-pn-realaudio"},
	Mime{".ram", "audio/x-pn-realaudio"},
	Mime{".rar", "application/octet-stream"},
	Mime{".ras", "image/x-cmu-raster"},
	Mime{".rc", "text/plain"},
	Mime{".rdf", "application/rdf+xml"},
	Mime{".rf", "image/vnd.rn-realflash"},
	Mime{".rgb", "image/x-rgb"},
	Mime{".rlf", "application/x-richlink"},
	Mime{".rm", "audio/x-pn-realaudio"},
	Mime{".rmf", "audio/x-rmf"},
	Mime{".rmi", "audio/mid"},
	Mime{".rmm", "audio/x-pn-realaudio"},
	Mime{".rmvb", "audio/x-pn-realaudio"},
	Mime{".rnx", "application/vnd.rn-realplayer"},
	Mime{".roff", "application/x-troff"},
	Mime{".rp", "image/vnd.rn-realpix"},
	Mime{".rpm", "audio/x-pn-realaudio-plugin"},
	Mime{".rt", "text/vnd.rn-realtext"},
	Mime{".rte", "x-lml/x-gps"},
	Mime{".rtf", "application/rtf"},
	Mime{".rtg", "application/metastream"},
	Mime{".rtx", "text/richtext"},
	Mime{".rv", "video/vnd.rn-realvideo"},
	Mime{".rwc", "application/x-rogerwilco"},
	Mime{".s3m", "audio/x-mod"},
	Mime{".s3z", "audio/x-mod"},
	Mime{".sca", "application/x-supercard"},
	Mime{".scd", "application/x-msschedule"},
	Mime{".sct", "text/scriptlet"},
	Mime{".sdf", "application/e-score"},
	Mime{".sea", "application/x-stuffit"},
	Mime{".setpay", "application/set-payment-initiation"},
	Mime{".setreg", "application/set-registration-initiation"},
	Mime{".sgm", "text/x-sgml"},
	Mime{".sgml", "text/x-sgml"},
	Mime{".sh", "application/x-sh"},
	Mime{".shar", "application/x-shar"},
	Mime{".shtml", "magnus-internal/parsed-html"},
	Mime{".shw", "application/presentations"},
	Mime{".si6", "image/si6"},
	Mime{".si7", "image/vnd.stiwap.sis"},
	Mime{".si9", "image/vnd.lgtwap.sis"},
	Mime{".sis", "application/vnd.symbian.install"},
	Mime{".sit", "application/x-stuffit"},
	Mime{".skd", "application/x-Koan"},
	Mime{".skm", "application/x-Koan"},
	Mime{".skp", "application/x-Koan"},
	Mime{".skt", "application/x-Koan"},
	Mime{".slc", "application/x-salsa"},
	Mime{".smd", "audio/x-smd"},
	Mime{".smi", "application/smil"},
	Mime{".smil", "application/smil"},
	Mime{".smp", "application/studiom"},
	Mime{".smz", "audio/x-smd"},
	Mime{".snd", "audio/basic"},
	Mime{".spc", "application/x-pkcs7-certificates"},
	Mime{".spl", "application/futuresplash"},
	Mime{".spr", "application/x-sprite"},
	Mime{".sprite", "application/x-sprite"},
	Mime{".sdp", "application/sdp"},
	Mime{".spt", "application/x-spt"},
	Mime{".src", "application/x-wais-source"},
	Mime{".sst", "application/vnd.ms-pkicertstore"},
	Mime{".stk", "application/hyperstudio"},
	Mime{".stl", "application/vnd.ms-pkistl"},
	Mime{".stm", "text/html"},
	Mime{".svg", "image/svg+xml"},
	Mime{".sv4cpio", "application/x-sv4cpio"},
	Mime{".sv4crc", "application/x-sv4crc"},
	Mime{".svf", "image/vnd"},
	Mime{".svg", "image/svg+xml"},
	Mime{".svh", "image/svh"},
	Mime{".svr", "x-world/x-svr"},
	Mime{".swf", "application/x-shockwave-flash"},
	Mime{".swfl", "application/x-shockwave-flash"},
	Mime{".t", "application/x-troff"},
	Mime{".tad", "application/octet-stream"},
	Mime{".talk", "text/x-speech"},
	Mime{".tar", "application/x-tar"},
	Mime{".taz", "application/x-tar"},
	Mime{".tbp", "application/x-timbuktu"},
	Mime{".tbt", "application/x-timbuktu"},
	Mime{".tcl", "application/x-tcl"},
	Mime{".tex", "application/x-tex"},
	Mime{".texi", "application/x-texinfo"},
	Mime{".texinfo", "application/x-texinfo"},
	Mime{".tgz", "application/x-compressed"},
	Mime{".thm", "application/vnd.eri.thm"},
	Mime{".tif", "image/tiff"},
	Mime{".tiff", "image/tiff"},
	Mime{".tki", "application/x-tkined"},
	Mime{".tkined", "application/x-tkined"},
	Mime{".toc", "application/toc"},
	Mime{".toy", "image/toy"},
	Mime{".tr", "application/x-troff"},
	Mime{".trk", "x-lml/x-gps"},
	Mime{".trm", "application/x-msterminal"},
	Mime{".tsi", "audio/tsplayer"},
	Mime{".tsp", "application/dsptype"},
	Mime{".tsv", "text/tab-separated-values"},
	Mime{".ttf", "application/octet-stream"},
	Mime{".ttz", "application/t-time"},
	Mime{".txt", "text/plain"},
	Mime{".uls", "text/iuls"},
	Mime{".ult", "audio/x-mod"},
	Mime{".ustar", "application/x-ustar"},
	Mime{".uu", "application/x-uuencode"},
	Mime{".uue", "application/x-uuencode"},
	Mime{".vcd", "application/x-cdlink"},
	Mime{".vcf", "text/x-vcard"},
	Mime{".vdo", "video/vdo"},
	Mime{".vib", "audio/vib"},
	Mime{".viv", "video/vivo"},
	Mime{".vivo", "video/vivo"},
	Mime{".vmd", "application/vocaltec-media-desc"},
	Mime{".vmf", "application/vocaltec-media-file"},
	Mime{".vmi", "application/x-dreamcast-vms-info"},
	Mime{".vms", "application/x-dreamcast-vms"},
	Mime{".vox", "audio/voxware"},
	Mime{".vqe", "audio/x-twinvq-plugin"},
	Mime{".vqf", "audio/x-twinvq"},
	Mime{".vql", "audio/x-twinvq"},
	Mime{".vre", "x-world/x-vream"},
	Mime{".vrml", "x-world/x-vrml"},
	Mime{".vrt", "x-world/x-vrt"},
	Mime{".vrw", "x-world/x-vream"},
	Mime{".vts", "workbook/formulaone"},
	Mime{".wav", "audio/x-wav"},
	Mime{".wax", "audio/x-ms-wax"},
	Mime{".wbmp", "image/vnd.wap.wbmp"},
	Mime{".wcm", "application/vnd.ms-works"},
	Mime{".wdb", "application/vnd.ms-works"},
	Mime{".web", "application/vnd.xara"},
	Mime{".wi", "image/wavelet"},
	Mime{".wis", "application/x-InstallShield"},
	Mime{".wks", "application/vnd.ms-works"},
	Mime{".wm", "video/x-ms-wm"},
	Mime{".wma", "audio/x-ms-wma"},
	Mime{".wmd", "application/x-ms-wmd"},
	Mime{".wmf", "application/x-msmetafile"},
	Mime{".wml", "text/vnd.wap.wml"},
	Mime{".wmlc", "application/vnd.wap.wmlc"},
	Mime{".wmls", "text/vnd.wap.wmlscript"},
	Mime{".wmlsc", "application/vnd.wap.wmlscriptc"},
	Mime{".wmlscript", "text/vnd.wap.wmlscript"},
	Mime{".wmv", "audio/x-ms-wmv"},
	Mime{".wmx", "video/x-ms-wmx"},
	Mime{".wmz", "application/x-ms-wmz"},
	Mime{".wpng", "image/x-up-wpng"},
	Mime{".wps", "application/vnd.ms-works"},
	Mime{".wpt", "x-lml/x-gps"},
	Mime{".wri", "application/x-mswrite"},
	Mime{".wrl", "x-world/x-vrml"},
	Mime{".wrz", "x-world/x-vrml"},
	Mime{".ws", "text/vnd.wap.wmlscript"},
	Mime{".wsc", "application/vnd.wap.wmlscriptc"},
	Mime{".wv", "video/wavelet"},
	Mime{".wvx", "video/x-ms-wvx"},
	Mime{".wxl", "application/x-wxl"},
	Mime{".x-gzip", "application/x-gzip"},
	Mime{".xaf", "x-world/x-vrml"},
	Mime{".xar", "application/vnd.xara"},
	Mime{".xbm", "image/x-xbitmap"},
	Mime{".xdm", "application/x-xdma"},
	Mime{".xdma", "application/x-xdma"},
	Mime{".xdw", "application/vnd.fujixerox.docuworks"},
	Mime{".xht", "application/xhtml+xml"},
	Mime{".xhtm", "application/xhtml+xml"},
	Mime{".xhtml", "application/xhtml+xml"},
	Mime{".xla", "application/vnd.ms-excel"},
	Mime{".xlc", "application/vnd.ms-excel"},
	Mime{".xll", "application/x-excel"},
	Mime{".xlm", "application/vnd.ms-excel"},
	Mime{".xls", "application/vnd.ms-excel"},
	Mime{".xlsx", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"},
	Mime{".xlt", "application/vnd.ms-excel"},
	Mime{".xlw", "application/vnd.ms-excel"},
	Mime{".xm", "audio/x-mod"},
	{".xml", "text/plain"},
	{".xml", "application/xml"},
	Mime{".xmz", "audio/x-mod"},
	Mime{".xof", "x-world/x-vrml"},
	Mime{".xpi", "application/x-xpinstall"},
	Mime{".xpm", "image/x-xpixmap"},
	Mime{".xsit", "text/xml"},
	Mime{".xsl", "text/xml"},
	Mime{".xul", "text/xul"},
	Mime{".xwd", "image/x-xwindowdump"},
	Mime{".xyz", "chemical/x-pdb"},
	Mime{".yz1", "application/x-yz1"},
	Mime{".z", "application/x-compress"},
	Mime{".zac", "application/x-zaurus-zac"},
	Mime{".zip", "application/zip"},
	Mime{".json", "application/json"},
}

var (
	UKNOWN       = Mime{"", ""}
	TEXT232      = Mime{".323", "text/h323"}
	VIDEO3GP     = Mime{".3gp", "video/3gpp"}
	APPAAB       = Mime{".aab", "application/x-authoware-bin"}
	APPAAM       = Mime{".aam", "application/x-authoware-map"}
	APPAAS       = Mime{".aas", "application/x-authoware-seg"}
	APPACX       = Mime{".acx", "application/internet-property-stream"}
	APPAI        = Mime{".ai", "application/postscript"}
	AUDIOAIF     = Mime{".aif", "audio/x-aiff"}
	AUDIOAIFC    = Mime{".aifc", "audio/x-aiff"}
	AUDIOAIFF    = Mime{".aiff", "audio/x-aiff"}
	AUDIOALS     = Mime{".als", "audio/X-Alpha5"}
	APPAMC       = Mime{".amc", "application/x-mpeg"}
	APPANI       = Mime{".ani", "application/octet-stream"}
	APPAPK       = Mime{".apk", "application/vnd.android.package-archive"}
	TEXTASC      = Mime{".asc", "text/plain"}
	APPASD       = Mime{".asd", "application/astound"}
	VIDEOASF     = Mime{".asf", "video/x-ms-asf"}
	APPASN       = Mime{".asn", "application/astound"}
	APPASP       = Mime{".asp", "application/x-asap"}
	VIDEOASR     = Mime{".asr", "video/x-ms-asf"}
	VIDEOASX     = Mime{".asx", "video/x-ms-asf"}
	AUDIOAU      = Mime{".au", "audio/basic"}
	APPAVB       = Mime{".avb", "application/octet-stream"}
	VIDEOAVI     = Mime{".avi", "video/x-msvideo"}
	AUDIOAWB     = Mime{".awb", "audio/amr-wb"}
	APPAXS       = Mime{".axs", "application/olescript"}
	TEXTBAS      = Mime{".bas", "text/plain"}
	APPBCPIO     = Mime{".bcpio", "application/x-bcpio"}
	APPBIN       = Mime{".bin ", "application/octet-stream"}
	APPBLD       = Mime{".bld", "application/bld"}
	APPBLD2      = Mime{".bld2", "application/bld2"}
	IMAGEBMP     = Mime{".bmp", "image/bmp"}
	APPBPK       = Mime{".bpk", "application/octet-stream"}
	APPBZ2       = Mime{".bz2", "application/x-bzip2"}
	TEXTC        = Mime{".c", "text/plain"}
	IMAGECAL     = Mime{".cal", "image/x-cals"}
	APPCAT       = Mime{".cat", "application/vnd.ms-pkiseccat"}
	APPCCN       = Mime{".ccn", "application/x-cnc"}
	APPCCO       = Mime{".cco", "application/x-cocoa"}
	APPCDF       = Mime{".cdf", "application/x-cdf"}
	APPCER       = Mime{".cer", "application/x-x509-ca-cert"}
	MAGINTCGI    = Mime{".cgi", "magnus-internal/cgi"}
	APPCHAT      = Mime{".chat", "application/x-chat"}
	APPCLASS     = Mime{".class", "application/octet-stream"}
	APPCLP       = Mime{".clp", "application/x-msclip"}
	IMAGECMX     = Mime{".cmx", "image/x-cmx"}
	APPCO        = Mime{".co", "application/x-cult3d-object"}
	IMAGECOD     = Mime{".cod", "image/cis-cod"}
	TEXTCONF     = Mime{".conf", "text/plain"}
	APPCPIO      = Mime{".cpio", "application/x-cpio"}
	TEXTCPP      = Mime{".cpp", "text/plain"}
	APPCPT       = Mime{".cpt", "application/mac-compactpro"}
	APPCRD       = Mime{".crd", "application/x-mscardfile"}
	APPCRL       = Mime{".crl", "application/pkix-crl"}
	APPCRT       = Mime{".crt", "application/x-x509-ca-cert"}
	APPCSH       = Mime{".csh", "application/x-csh"}
	CHEMICALCSM  = Mime{".csm", "chemical/x-csml"}
	CHEMICALCSML = Mime{".csml", "chemical/x-csml"}
	TEXTCSS      = Mime{".css", "text/css"}
	APPCUR       = Mime{".cur", "application/octet-stream"}
	XLMLDCM      = Mime{".dcm", "x-lml/x-evm"}
	APPDCR       = Mime{".dcr", "application/x-director"}
	IMAGEDCX     = Mime{".dcx", "image/x-dcx"}
	APPDER       = Mime{".der", "application/x-x509-ca-cert"}
	TEXTDHTML    = Mime{".dhtml", "text/html"}
	APPDIR       = Mime{".dir", "application/x-director"}
	APPDLL       = Mime{".dll", "application/x-msdownload"}
	APPDMG       = Mime{".dmg", "application/octet-stream"}
	APPDMS       = Mime{".dms", "application/octet-stream"}
	APPDOC       = Mime{".doc", "application/msword"}
	APPDOCX      = Mime{".docx", "application/vnd.openxmlformats-officedocument.wordprocessingml.document"}
	APPDOT       = Mime{".dot", "application/msword"}
	APPDVI       = Mime{".dvi", "application/x-dvi"}
	DRAWDWF      = Mime{".dwf", "drawing/x-dwf"}
	APPDWG       = Mime{".dwg", "application/x-autocad"}
	APPDXF       = Mime{".dxf", "application/x-autocad"}
	APPDXR       = Mime{".dxr", "application/x-director"}
	APPEBK       = Mime{".ebk", "application/x-expandedbook"}
	CHEMICALEMB  = Mime{".emb", "chemical/x-embl-dl-nucleotide"}
	ECEMICALEMBL = Mime{".embl", "chemical/x-embl-dl-nucleotide"}
	APPEPS       = Mime{".eps", "application/postscript"}
	APPEPUB      = Mime{".epub", "application/epub+zip"}
	IMAGEERI     = Mime{".eri", "image/x-eri"}
	AUDIOES      = Mime{".es", "audio/echospeech"}
	AUDIOESL     = Mime{".esl", "audio/echospeech"}
	APPETC       = Mime{".etc", "application/x-earthtime"}
	TEXTETX      = Mime{".etx", "text/x-setext"}
	XLMLEVM      = Mime{".evm", "x-lml/x-evm"}
	APPEVY       = Mime{".evy", "application/envoy"}
	APPEXE       = Mime{".exe", "application/octet-stream"}
	IMAGEFH4     = Mime{".fh4", "image/x-freehand"}
	IMAGEFH5     = Mime{".fh5", "image/x-freehand"}
	IMAGEFHC     = Mime{".fhc", "image/x-freehand"}
	APPFIF       = Mime{".fif", "application/fractals"}
	XWORLDFLR    = Mime{".flr", "x-world/x-vrml"}
	FLVAPPFLV    = Mime{".flv", "flv-application/octet-stream"}
	APPFM        = Mime{".fm", "application/x-maker"}
	IMAGEFPX     = Mime{".fpx", "image/x-fpx"}
	VIDEOFVI     = Mime{".fvi", "video/isivideo"}
	CHEMICALGAU  = Mime{".gau", "chemical/x-gaussian-input"}
	APPGCA       = Mime{".gca", "application/x-gca-compressed"}
	// = Mime{ ".gdb", "x-lml/x-gdb" }
	IMAGEGIF = Mime{".gif", "image/gif"}
	APPGPS   = Mime{".gps", "application/x-gps"}
	APPGTAR  = Mime{".gtar", "application/x-gtar"}
	APPGZ    = Mime{".gz", "application/x-gzip"}
	TEXTH    = Mime{".h", "text/plain"}
	// = Mime{ ".hdf", "application/x-hdf" }
	TEXTHDM  = Mime{".hdm", "text/x-hdml"}
	TEXTHDML = Mime{".hdml", "text/x-hdml"}
	// = Mime{ ".hlp", "application/winhlp" }
	// = Mime{ ".hqx", "application/mac-binhex40" }
	// = Mime{ ".hta", "application/hta" }
	// = Mime{ ".htc", "text/x-component" }
	TEXTHTM  = Mime{".htm", "text/html"}
	TEXTHTML = Mime{".html", "text/html"}
	// = Mime{ ".hts", "text/html" }
	// = Mime{ ".htt", "text/webviewhtml" }
	// = Mime{ ".ice", "x-conference/x-cooltalk" }
	// = Mime{ ".ico", "image/x-icon" }
	// = Mime{ ".ief", "image/ief" }
	// = Mime{ ".ifm", "image/gif" }
	// = Mime{ ".ifs", "image/ifs" }
	// = Mime{ ".iii", "application/x-iphone" }
	// = Mime{ ".imy", "audio/melody" }
	// = Mime{ ".ins", "application/x-internet-signup" }
	// = Mime{ ".ips", "application/x-ipscript" }
	// = Mime{ ".ipx", "application/x-ipix" }
	// = Mime{ ".isp", "application/x-internet-signup" }
	// = Mime{ ".it", "audio/x-mod" }
	// = Mime{ ".itz", "audio/x-mod" }
	// = Mime{ ".ivr", "i-world/i-vrml" }
	IMAGEJ2K = Mime{".j2k", "image/j2k"}
	// = Mime{ ".jad", "text/vnd.sun.j2me.app-descriptor" }
	// = Mime{ ".jam", "application/x-jam" }
	APPJAR = Mime{".jar", "application/java-archive"}
	// = Mime{ ".java", "text/plain" }
	// = Mime{ ".jfif", "image/pipeg" }
	// = Mime{ ".jnlp", "application/x-java-jnlp-file" }
	IMAGEJPE  = Mime{".jpe", "image/jpeg"}
	IMAGEJPEG = Mime{".jpeg", "image/jpeg"}
	IMAGEJPG  = Mime{".jpg", "image/jpeg"}
	IMAGEJPZ  = Mime{".jpz", "image/jpeg"}
	APPJS     = Mime{".js", "application/x-javascript"}
	// = Mime{ ".jwc", "application/jwc" }
	// = Mime{ ".kjx", "application/x-kjx" }
	// = Mime{ ".lak", "x-lml/x-lak" }
	// = Mime{ ".latex", "application/x-latex" }
	// = Mime{ ".lcc", "application/fastman" }
	// = Mime{ ".lcl", "application/x-digitalloca" }
	// = Mime{ ".lcr", "application/x-digitalloca" }
	// = Mime{ ".lgh", "application/lgh" }
	// = Mime{ ".lha", "application/octet-stream" }
	// = Mime{ ".lml", "x-lml/x-lml" }
	// = Mime{ ".lmlpack", "x-lml/x-lmlpack" }
	// = Mime{ ".log", "text/plain" }
	// = Mime{ ".lsf", "video/x-la-asf" }
	// = Mime{ ".lsx", "video/x-la-asf" }
	// = Mime{ ".lzh", "application/octet-stream" }
	// = Mime{ ".m13", "application/x-msmediaview" }
	// = Mime{ ".m14", "application/x-msmediaview" }
	// = Mime{ ".m15", "audio/x-mod" }
	// = Mime{ ".m3u", "audio/x-mpegurl" }
	// = Mime{ ".m3url", "audio/x-mpegurl" }
	// = Mime{ ".m4a", "audio/mp4a-latm" }
	// = Mime{ ".m4b", "audio/mp4a-latm" }
	// = Mime{ ".m4p", "audio/mp4a-latm" }
	VIDEOM4U = Mime{".m4u", "video/vnd.mpegurl"}
	VIDEOM4V = Mime{".m4v", "video/x-m4v"}
	// = Mime{ ".ma1", "audio/ma1" }
	// = Mime{ ".ma2", "audio/ma2" }
	// = Mime{ ".ma3", "audio/ma3" }
	// = Mime{ ".ma5", "audio/ma5" }
	// = Mime{ ".man", "application/x-troff-man" }
	// = Mime{ ".map", "magnus-internal/imagemap" }
	// = Mime{ ".mbd", "application/mbedlet" }
	// = Mime{ ".mct", "application/x-mascot" }
	// = Mime{ ".mdb", "application/x-msaccess" }
	// = Mime{ ".mdz", "audio/x-mod" }
	// = Mime{ ".me", "application/x-troff-me" }
	// = Mime{ ".mel", "text/x-vmel" }
	// = Mime{ ".mht", "message/rfc822" }
	// = Mime{ ".mhtml", "message/rfc822" }
	// = Mime{ ".mi", "application/x-mif" }
	AUDIOMID  = Mime{".mid", "audio/mid"}
	AUDIOMIDI = Mime{".midi", "audio/midi"}
	// = Mime{ ".mif", "application/x-mif" }
	// = Mime{ ".mil", "image/x-cals" }
	// = Mime{ ".mio", "audio/x-mio" }
	// = Mime{ ".mmf", "application/x-skt-lbs" }
	// = Mime{ ".mng", "video/x-mng" }
	// = Mime{ ".mny", "application/x-msmoney" }
	// = Mime{ ".moc", "application/x-mocha" }
	// = Mime{ ".mocha", "application/x-mocha" }
	// = Mime{ ".mod", "audio/x-mod" }
	// = Mime{ ".mof", "application/x-yumekara" }
	// = Mime{ ".mol", "chemical/x-mdl-molfile" }
	// = Mime{ ".mop", "chemical/x-mopac-input" }
	VIDEOMOV   = Mime{".mov", "video/quicktime"}
	VIDEOMOVIE = Mime{".movie", "video/x-sgi-movie"}
	VIDEOMP2   = Mime{".mp2", "video/mpeg"}
	AUDIOMP3   = Mime{".mp3", "audio/mpeg"}
	VIDEOMP4   = Mime{".mp4", "video/mp4"}
	VIDEOMPA   = Mime{".mpa", "video/mpeg"}
	// = Mime{ ".mpc", "application/vnd.mpohun.certificate" }
	VIDEOMPE  = Mime{".mpe", "video/mpeg"}
	VIDEOMPEG = Mime{".mpeg", "video/mpeg"}
	VIDEOMPG  = Mime{".mpg", "video/mpeg"}
	VIDEOMPG4 = Mime{".mpg4", "video/mp4"}
	AUDIOMPGA = Mime{".mpga", "audio/mpeg"}
	// = Mime{ ".mpn", "application/vnd.mophun.application" }
	// = Mime{ ".mpp", "application/vnd.ms-project" }
	// = Mime{ ".mps", "application/x-mapserver" }
	// = Mime{ ".mpv2", "video/mpeg" }
	// = Mime{ ".mrl", "text/x-mrml" }
	// = Mime{ ".mrm", "application/x-mrm" }
	// = Mime{ ".ms", "application/x-troff-ms" }
	// = Mime{ ".msg", "application/vnd.ms-outlook" }
	// = Mime{ ".mts", "application/metastream" }
	// = Mime{ ".mtx", "application/metastream" }
	// = Mime{ ".mtz", "application/metastream" }
	// = Mime{ ".mvb", "application/x-msmediaview" }
	// = Mime{ ".mzv", "application/metastream" }
	// = Mime{ ".nar", "application/zip" }
	IMAGENBMP = Mime{".nbmp", "image/nbmp"}
	// = Mime{ ".nc", "application/x-netcdf" }
	// = Mime{ ".ndb", "x-lml/x-ndb" }
	// = Mime{ ".ndwn", "application/ndwn" }
	// = Mime{ ".nif", "application/x-nif" }
	// = Mime{ ".nmz", "application/x-scream" }
	// = Mime{ ".nokia-op-logo", "image/vnd.nok-oplogo-color" }
	// = Mime{ ".npx", "application/x-netfpx" }
	// = Mime{ ".nsnd", "audio/nsnd" }
	// = Mime{ ".nva", "application/x-neva1" }
	// = Mime{ ".nws", "message/rfc822" }
	// = Mime{ ".oda", "application/oda" }
	AUDIOOGG = Mime{".ogg", "audio/ogg"}
	// = Mime{ ".oom", "application/x-AtlasMate-Plugin" }
	// = Mime{ ".p10", "application/pkcs10" }
	// = Mime{ ".p12", "application/x-pkcs12" }
	// = Mime{ ".p7b", "application/x-pkcs7-certificates" }
	// = Mime{ ".p7c", "application/x-pkcs7-mime" }
	// = Mime{ ".p7m", "application/x-pkcs7-mime" }
	// = Mime{ ".p7r", "application/x-pkcs7-certreqresp" }
	// = Mime{ ".p7s", "application/x-pkcs7-signature" }
	// = Mime{ ".pac", "audio/x-pac" }
	// = Mime{ ".pae", "audio/x-epac" }
	// = Mime{ ".pan", "application/x-pan" }
	// = Mime{ ".pbm", "image/x-portable-bitmap" }
	// = Mime{ ".pcx", "image/x-pcx" }
	// = Mime{ ".pda", "image/x-pda" }
	// = Mime{ ".pdb", "chemical/x-pdb" }
	// = Mime{ ".pdf", "application/pdf" }
	// = Mime{ ".pfr", "application/font-tdpfr" }
	// = Mime{ ".pfx", "application/x-pkcs12" }
	// = Mime{ ".pgm", "image/x-portable-graymap" }
	// = Mime{ ".pict", "image/x-pict" }
	// = Mime{ ".pko", "application/ynd.ms-pkipko" }
	// = Mime{ ".pm", "application/x-perl" }
	// = Mime{ ".pma", "application/x-perfmon" }
	// = Mime{ ".pmc", "application/x-perfmon" }
	// = Mime{ ".pmd", "application/x-pmd" }
	// = Mime{ ".pml", "application/x-perfmon" }
	// = Mime{ ".pmr", "application/x-perfmon" }
	// = Mime{ ".pmw", "application/x-perfmon" }
	IMAGEPNG = Mime{".png", "image/png"}
	// = Mime{ ".pnm", "image/x-portable-anymap" }
	// = Mime{ ".pnz", "image/png" }
	// = Mime{ ".pot,", "application/vnd.ms-powerpoint" }
	// = Mime{ ".ppm", "image/x-portable-pixmap" }
	// = Mime{ ".pps", "application/vnd.ms-powerpoint" }
	// = Mime{ ".ppt", "application/vnd.ms-powerpoint" }
	// = Mime{ ".pptx", "application/vnd.openxmlformats-officedocument.presentationml.presentation" }
	// = Mime{ ".pqf", "application/x-cprplayer" }
	// = Mime{ ".pqi", "application/cprplayer" }
	// = Mime{ ".prc", "application/x-prc" }
	// = Mime{ ".prf", "application/pics-rules" }
	// = Mime{ ".prop", "text/plain" }
	// = Mime{ ".proxy", "application/x-ns-proxy-autoconfig" }
	// = Mime{ ".ps", "application/postscript" }
	// = Mime{ ".ptlk", "application/listenup" }
	// = Mime{ ".pub", "application/x-mspublisher" }
	// = Mime{ ".pvx", "video/x-pv-pvx" }
	// = Mime{ ".qcp", "audio/vnd.qcelp" }
	// = Mime{ ".qt", "video/quicktime" }
	// = Mime{ ".qti", "image/x-quicktime" }
	// = Mime{ ".qtif", "image/x-quicktime" }
	// = Mime{ ".r3t", "text/vnd.rn-realtext3d" }
	// = Mime{ ".ra", "audio/x-pn-realaudio" }
	// = Mime{ ".ram", "audio/x-pn-realaudio" }
	APPRAR = Mime{".rar", "application/octet-stream"}
	// = Mime{ ".ras", "image/x-cmu-raster" }
	// = Mime{ ".rc", "text/plain" }
	// = Mime{ ".rdf", "application/rdf+xml" }
	// = Mime{ ".rf", "image/vnd.rn-realflash" }
	// = Mime{ ".rgb", "image/x-rgb" }
	// = Mime{ ".rlf", "application/x-richlink" }
	// = Mime{ ".rm", "audio/x-pn-realaudio" }
	// = Mime{ ".rmf", "audio/x-rmf" }
	// = Mime{ ".rmi", "audio/mid" }
	// = Mime{ ".rmm", "audio/x-pn-realaudio" }
	// = Mime{ ".rmvb", "audio/x-pn-realaudio" }
	// = Mime{ ".rnx", "application/vnd.rn-realplayer" }
	// = Mime{ ".roff", "application/x-troff" }
	// = Mime{ ".rp", "image/vnd.rn-realpix" }
	// = Mime{ ".rpm", "audio/x-pn-realaudio-plugin" }
	// = Mime{ ".rt", "text/vnd.rn-realtext" }
	// = Mime{ ".rte", "x-lml/x-gps" }
	// = Mime{ ".rtf", "application/rtf" }
	// = Mime{ ".rtg", "application/metastream" }
	// = Mime{ ".rtx", "text/richtext" }
	// = Mime{ ".rv", "video/vnd.rn-realvideo" }
	// = Mime{ ".rwc", "application/x-rogerwilco" }
	// = Mime{ ".s3m", "audio/x-mod" }
	// = Mime{ ".s3z", "audio/x-mod" }
	// = Mime{ ".sca", "application/x-supercard" }
	// = Mime{ ".scd", "application/x-msschedule" }
	// = Mime{ ".sct", "text/scriptlet" }
	// = Mime{ ".sdf", "application/e-score" }
	// = Mime{ ".sea", "application/x-stuffit" }
	// = Mime{ ".setpay", "application/set-payment-initiation" }
	// = Mime{ ".setreg", "application/set-registration-initiation" }
	// = Mime{ ".sgm", "text/x-sgml" }
	// = Mime{ ".sgml", "text/x-sgml" }
	// = Mime{ ".sh", "application/x-sh" }
	// = Mime{ ".shar", "application/x-shar" }
	// = Mime{ ".shtml", "magnus-internal/parsed-html" }
	// = Mime{ ".shw", "application/presentations" }
	// = Mime{ ".si6", "image/si6" }
	// = Mime{ ".si7", "image/vnd.stiwap.sis" }
	// = Mime{ ".si9", "image/vnd.lgtwap.sis" }
	// = Mime{ ".sis", "application/vnd.symbian.install" }
	// = Mime{ ".sit", "application/x-stuffit" }
	// = Mime{ ".skd", "application/x-Koan" }
	// = Mime{ ".skm", "application/x-Koan" }
	// = Mime{ ".skp", "application/x-Koan" }
	// = Mime{ ".skt", "application/x-Koan" }
	// = Mime{ ".slc", "application/x-salsa" }
	// = Mime{ ".smd", "audio/x-smd" }
	// = Mime{ ".smi", "application/smil" }
	// = Mime{ ".smil", "application/smil" }
	// = Mime{ ".smp", "application/studiom" }
	// = Mime{ ".smz", "audio/x-smd" }
	// = Mime{ ".snd", "audio/basic" }
	// = Mime{ ".spc", "application/x-pkcs7-certificates" }
	// = Mime{ ".spl", "application/futuresplash" }
	// = Mime{ ".spr", "application/x-sprite" }
	// = Mime{ ".sprite", "application/x-sprite" }
	// = Mime{ ".sdp", "application/sdp" }
	// = Mime{ ".spt", "application/x-spt" }
	// = Mime{ ".src", "application/x-wais-source" }
	// = Mime{ ".sst", "application/vnd.ms-pkicertstore" }
	// = Mime{ ".stk", "application/hyperstudio" }
	// = Mime{ ".stl", "application/vnd.ms-pkistl" }
	// = Mime{ ".stm", "text/html" }
	// = Mime{ ".sv4cpio", "application/x-sv4cpio" }
	// = Mime{ ".sv4crc", "application/x-sv4crc" }
	IMAGESVF = Mime{".svf", "image/vnd"}
	IMAGESVG = Mime{".svg", "image/svg+xml"}
	IMAGESVH = Mime{".svh", "image/svh"}
	// = Mime{ ".svr", "x-world/x-svr" }
	APPSWF  = Mime{".swf", "application/x-shockwave-flash"}
	APPSWFL = Mime{".swfl", "application/x-shockwave-flash"}
	// = Mime{ ".t", "application/x-troff" }
	// = Mime{ ".tad", "application/octet-stream" }
	// = Mime{ ".talk", "text/x-speech" }
	APPTAR = Mime{".tar", "application/x-tar"}
	APPTAZ = Mime{".taz", "application/x-tar"}
	// = Mime{ ".tbp", "application/x-timbuktu" }
	// = Mime{ ".tbt", "application/x-timbuktu" }
	// = Mime{ ".tcl", "application/x-tcl" }
	// = Mime{ ".tex", "application/x-tex" }
	// = Mime{ ".texi", "application/x-texinfo" }
	// = Mime{ ".texinfo", "application/x-texinfo" }
	// = Mime{ ".tgz", "application/x-compressed" }
	// = Mime{ ".thm", "application/vnd.eri.thm" }
	// = Mime{ ".tif", "image/tiff" }
	// = Mime{ ".tiff", "image/tiff" }
	// = Mime{ ".tki", "application/x-tkined" }
	// = Mime{ ".tkined", "application/x-tkined" }
	// = Mime{ ".toc", "application/toc" }
	// = Mime{ ".toy", "image/toy" }
	// = Mime{ ".tr", "application/x-troff" }
	// = Mime{ ".trk", "x-lml/x-gps" }
	// = Mime{ ".trm", "application/x-msterminal" }
	// = Mime{ ".tsi", "audio/tsplayer" }
	// = Mime{ ".tsp", "application/dsptype" }
	// = Mime{ ".tsv", "text/tab-separated-values" }
	APPTTF = Mime{".ttf", "application/octet-stream"}
	// = Mime{ ".ttz", "application/t-time" }
	TEXTTXT = Mime{".txt", "text/plain"}
	// = Mime{ ".uls", "text/iuls" }
	// = Mime{ ".ult", "audio/x-mod" }
	// = Mime{ ".ustar", "application/x-ustar" }
	// = Mime{ ".uu", "application/x-uuencode" }
	// = Mime{ ".uue", "application/x-uuencode" }
	// = Mime{ ".vcd", "application/x-cdlink" }
	// = Mime{ ".vcf", "text/x-vcard" }
	// = Mime{ ".vdo", "video/vdo" }
	// = Mime{ ".vib", "audio/vib" }
	// = Mime{ ".viv", "video/vivo" }
	// = Mime{ ".vivo", "video/vivo" }
	// = Mime{ ".vmd", "application/vocaltec-media-desc" }
	// = Mime{ ".vmf", "application/vocaltec-media-file" }
	// = Mime{ ".vmi", "application/x-dreamcast-vms-info" }
	// = Mime{ ".vms", "application/x-dreamcast-vms" }
	// = Mime{ ".vox", "audio/voxware" }
	// = Mime{ ".vqe", "audio/x-twinvq-plugin" }
	// = Mime{ ".vqf", "audio/x-twinvq" }
	// = Mime{ ".vql", "audio/x-twinvq" }
	// = Mime{ ".vre", "x-world/x-vream" }
	// = Mime{ ".vrml", "x-world/x-vrml" }
	// = Mime{ ".vrt", "x-world/x-vrt" }
	// = Mime{ ".vrw", "x-world/x-vream" }
	// = Mime{ ".vts", "workbook/formulaone" }
	// = Mime{ ".wav", "audio/x-wav" }
	// = Mime{ ".wax", "audio/x-ms-wax" }
	// = Mime{ ".wbmp", "image/vnd.wap.wbmp" }
	// = Mime{ ".wcm", "application/vnd.ms-works" }
	// = Mime{ ".wdb", "application/vnd.ms-works" }
	// = Mime{ ".web", "application/vnd.xara" }
	// = Mime{ ".wi", "image/wavelet" }
	// = Mime{ ".wis", "application/x-InstallShield" }
	// = Mime{ ".wks", "application/vnd.ms-works" }
	// = Mime{ ".wm", "video/x-ms-wm" }
	AUDIOWMA = Mime{".wma", "audio/x-ms-wma"}
	// = Mime{ ".wmd", "application/x-ms-wmd" }
	// = Mime{ ".wmf", "application/x-msmetafile" }
	// = Mime{ ".wml", "text/vnd.wap.wml" }
	// = Mime{ ".wmlc", "application/vnd.wap.wmlc" }
	// = Mime{ ".wmls", "text/vnd.wap.wmlscript" }
	// = Mime{ ".wmlsc", "application/vnd.wap.wmlscriptc" }
	// = Mime{ ".wmlscript", "text/vnd.wap.wmlscript" }
	// = Mime{ ".wmv", "audio/x-ms-wmv" }
	// = Mime{ ".wmx", "video/x-ms-wmx" }
	// = Mime{ ".wmz", "application/x-ms-wmz" }
	// = Mime{ ".wpng", "image/x-up-wpng" }
	// = Mime{ ".wps", "application/vnd.ms-works" }
	// = Mime{ ".wpt", "x-lml/x-gps" }
	// = Mime{ ".wri", "application/x-mswrite" }
	// = Mime{ ".wrl", "x-world/x-vrml" }
	// = Mime{ ".wrz", "x-world/x-vrml" }
	// = Mime{ ".ws", "text/vnd.wap.wmlscript" }
	// = Mime{ ".wsc", "application/vnd.wap.wmlscriptc" }
	// = Mime{ ".wv", "video/wavelet" }
	// = Mime{ ".wvx", "video/x-ms-wvx" }
	// = Mime{ ".wxl", "application/x-wxl" }
	// = Mime{ ".x-gzip", "application/x-gzip" }
	// = Mime{ ".xaf", "x-world/x-vrml" }
	// = Mime{ ".xar", "application/vnd.xara" }
	// = Mime{ ".xbm", "image/x-xbitmap" }
	// = Mime{ ".xdm", "application/x-xdma" }
	// = Mime{ ".xdma", "application/x-xdma" }
	// = Mime{ ".xdw", "application/vnd.fujixerox.docuworks" }
	// = Mime{ ".xht", "application/xhtml+xml" }
	// = Mime{ ".xhtm", "application/xhtml+xml" }
	// = Mime{ ".xhtml", "application/xhtml+xml" }
	// = Mime{ ".xla", "application/vnd.ms-excel" }
	// = Mime{ ".xlc", "application/vnd.ms-excel" }
	// = Mime{ ".xll", "application/x-excel" }
	// = Mime{ ".xlm", "application/vnd.ms-excel" }
	APPXLS  = Mime{".xls", "application/vnd.ms-excel"}
	APPXLSX = Mime{".xlsx", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"}
	// = Mime{ ".xlt", "application/vnd.ms-excel" }
	// = Mime{ ".xlw", "application/vnd.ms-excel" }
	// = Mime{ ".xm", "audio/x-mod" }
	TEXTXML = Mime{".xml", "text/plain"}
	APPXML  = Mime{".xml", "application/xml"}
	// = Mime{ ".xmz", "audio/x-mod" }
	// = Mime{ ".xof", "x-world/x-vrml" }
	// = Mime{ ".xpi", "application/x-xpinstall" }
	// = Mime{ ".xpm", "image/x-xpixmap" }
	// = Mime{ ".xsit", "text/xml" }
	// = Mime{ ".xsl", "text/xml" }
	// = Mime{ ".xul", "text/xul" }
	// = Mime{ ".xwd", "image/x-xwindowdump" }
	// = Mime{ ".xyz", "chemical/x-pdb" }
	// = Mime{ ".yz1", "application/x-yz1" }
	// = Mime{ ".z", "application/x-compress" }
	// = Mime{ ".zac", "application/x-zaurus-zac" }
	APPZIP  = Mime{".zip", "application/zip"}
	APPJSON = Mime{".json", "application/json"}
)
