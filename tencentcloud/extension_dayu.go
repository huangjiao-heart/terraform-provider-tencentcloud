package tencentcloud

const (
	DAYU_RESOURCE_TYPE_BGPIP  = "bgpip"
	DAYU_RESOURCE_TYPE_BGP    = "bgp"
	DAYU_RESOURCE_TYPE_BGPMUL = "bgp-multip"
	DAYU_RESOURCE_TYPE_NET    = "net"
)

var DAYU_RESOURCE_TYPE = []string{DAYU_RESOURCE_TYPE_BGPIP, DAYU_RESOURCE_TYPE_BGP, DAYU_RESOURCE_TYPE_BGPMUL, DAYU_RESOURCE_TYPE_NET}

var DAYU_RESOURCE_TYPE_RULE = []string{DAYU_RESOURCE_TYPE_BGPIP, DAYU_RESOURCE_TYPE_NET}

var DAYU_RESOURCE_TYPE_CASE = []string{DAYU_RESOURCE_TYPE_BGPIP, DAYU_RESOURCE_TYPE_BGPMUL, DAYU_RESOURCE_TYPE_BGP}

const (
	DAYU_APP_PLATFORM_PC     = "PC"
	DAYU_APP_PLATFORM_MOBILE = "MOBILE"
	DAYU_APP_PLATFORM_TV     = "TV"
	DAYU_APP_PLATFORM_SERVER = "SERVER"
)

var DAYU_APP_PLATFORM = []string{DAYU_APP_PLATFORM_PC, DAYU_APP_PLATFORM_MOBILE, DAYU_APP_PLATFORM_TV, DAYU_APP_PLATFORM_SERVER}

const (
	DAYU_PROTOCOL_TCP  = "tcp"
	DAYU_PROTOCOL_UDP  = "udp"
	DAYU_PROTOCOL_ICMP = "icmp"
	DAYU_PROTOCOL_ALL  = "all"
)

var DAYU_PROTOCOL = []string{DAYU_PROTOCOL_TCP, DAYU_PROTOCOL_UDP, DAYU_PROTOCOL_ICMP, DAYU_PROTOCOL_ALL}

const (
	DAYU_MATCH_TYPE_SUNDAY = "sunday"
	DAYU_MATCH_TYPE_PCRE   = "pcre"
)

var DAYU_MATCH_TYPE = []string{DAYU_MATCH_TYPE_SUNDAY, DAYU_MATCH_TYPE_PCRE}

const (
	DAYU_MATCH_SWITCH_ON_L5 = "begin_l5"
	DAYU_MATCH_SWITCH_OFF   = "no_match"
)

var DAYU_MATCH_SWITCH = []string{DAYU_MATCH_SWITCH_ON_L5, DAYU_MATCH_SWITCH_OFF}

const (
	DAYU_IP_TYPE_BLACK = "black"
	DAYU_IP_TYPE_WHITE = "white"
)

var DAYU_IP_TYPE = []string{DAYU_IP_TYPE_BLACK, DAYU_IP_TYPE_WHITE}

const (
	DAYU_PACKET_ACTION_DROP             = "drop"
	DAYU_PACKET_ACTION_TRANSMIT         = "transmit"
	DAYU_PACKET_PROTOCOL_DROP_BLACK     = "drop_black"
	DAYU_PACKET_PROTOCOL_DROP_RST       = "drop_rst"
	DAYU_PACKET_PROTOCOL_DROP_BLACK_RST = "drop_black_rst"
)

var DAYU_PACKET_ACTION = []string{DAYU_PACKET_ACTION_DROP, DAYU_PACKET_ACTION_TRANSMIT, DAYU_PACKET_PROTOCOL_DROP_BLACK, DAYU_PACKET_PROTOCOL_DROP_RST, DAYU_PACKET_PROTOCOL_DROP_BLACK_RST}

const (
	DAYU_PORT_ACTION_DROP     = "drop"
	DAYU_PORT_ACTION_TRANSMIT = "transmit"
)

var DAYU_PORT_ACTION = []string{DAYU_PORT_ACTION_DROP, DAYU_PORT_ACTION_TRANSMIT}

const (
	DAYU_APP_TYPE_WEB   = "WEB"
	DAYU_APP_TYPE_GAME  = "GAME"
	DAYU_APP_TYPE_APP   = "APP"
	DAYU_APP_TYPE_OTHER = "OTHER"
)

var DAYU_APP_TYPE = []string{DAYU_APP_TYPE_WEB, DAYU_APP_TYPE_GAME, DAYU_APP_TYPE_APP, DAYU_APP_TYPE_OTHER}

const (
	DAYU_BOOL_FLAG_TRUE  = "yes"
	DAYU_BOOL_FLAG_FALSE = "no"
)

var DAYU_BOOL_FLAG = []string{DAYU_BOOL_FLAG_FALSE, DAYU_BOOL_FLAG_TRUE}

const (
	DAYU_L7_RULE_PROTOCOL_HTTP  = "http"
	DAYU_L7_RULE_PROTOCOL_HTTPS = "https"
)

var DAYU_L7_RULE_PROTOCOL = []string{DAYU_L7_RULE_PROTOCOL_HTTP, DAYU_L7_RULE_PROTOCOL_HTTPS}

const (
	DAYU_RULE_METHOD_GET  = "GET"
	DAYU_RULE_METHOD_HEAD = "HEAD"
)

var DAYU_RULE_METHOD = []string{
	DAYU_RULE_METHOD_GET,
	DAYU_RULE_METHOD_HEAD,
}

const (
	DAYU_L7_RULE_SOURCE_TYPE_HOST = 1
	DAYU_L7_RULE_SOURCE_TYPE_IP   = 2
)

var DAYU_L7_RULE_SOURCE_TYPE = []int{DAYU_L7_RULE_SOURCE_TYPE_HOST, DAYU_L7_RULE_SOURCE_TYPE_IP}

const (
	DAYU_L7_HTTPS_SWITCH_ON_DEFAULT = 20000
	DAYU_L7_HTTPS_SWITCH_OFF        = 0
)

const (
	DAYU_CC_POLICY_ACTION_DROP = "drop"
	DAYU_CC_POLICY_ACTION_ALG  = "alg"
)

var DAYU_CC_POLICY_ACTION = []string{DAYU_CC_POLICY_ACTION_DROP, DAYU_CC_POLICY_ACTION_ALG}

const (
	DAYU_CC_POLICY_SMODE_MATCH       = "matching"
	DAYU_CC_POLICY_SMODE_SPEED_LIMIT = "speedlimit"
)

var DAYU_CC_POLICY_SMODE = []string{DAYU_CC_POLICY_SMODE_MATCH, DAYU_CC_POLICY_SMODE_SPEED_LIMIT}

const (
	DAYU_CC_POLICY_CHECK_TYPE_HOST    = "host"
	DAYU_CC_POLICY_CHECK_TYPE_CGI     = "cgi"
	DAYU_CC_POLICY_CHECK_TYPE_UA      = "ua"
	DAYU_CC_POLICY_CHECK_TYPE_REFERER = "referer"
)

var DAYU_CC_POLICY_HTTP_CHECK_TYPE = []string{DAYU_CC_POLICY_CHECK_TYPE_HOST, DAYU_CC_POLICY_CHECK_TYPE_CGI, DAYU_CC_POLICY_CHECK_TYPE_UA, DAYU_CC_POLICY_CHECK_TYPE_REFERER}

var DAYU_CC_POLICY_HTTPS_CHECK_TYPE = []string{DAYU_CC_POLICY_CHECK_TYPE_CGI, DAYU_CC_POLICY_CHECK_TYPE_UA, DAYU_CC_POLICY_CHECK_TYPE_REFERER}

const (
	DAYU_CC_POLICY_CHECK_OP_INCLUDE     = "include"
	DAYU_CC_POLICY_CHECK_OP_NOT_INCLUDE = "not_include"
	DAYU_CC_POLICY_CHECK_OP_EQUAL       = "equal"
)

var DAYU_CC_POLICY_CHECK_OP = []string{DAYU_CC_POLICY_CHECK_OP_INCLUDE, DAYU_CC_POLICY_CHECK_OP_NOT_INCLUDE, DAYU_CC_POLICY_CHECK_OP_EQUAL}
var DAYU_CC_POLICY_CHECK_OP_HTTPS = []string{DAYU_CC_POLICY_CHECK_OP_INCLUDE, DAYU_CC_POLICY_CHECK_OP_EQUAL}

const (
	DAYU_L4_RULE_PROTOCOL_TCP = "TCP"
	DAYU_L4_RULE_PROTOCOL_UDP = "UDP"
)

var DAYU_L4_RULE_PROTOCOL = []string{DAYU_L4_RULE_PROTOCOL_TCP, DAYU_L4_RULE_PROTOCOL_UDP}
