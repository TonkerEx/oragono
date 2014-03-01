package irc

import (
	"errors"
	"regexp"
	"time"
)

var (
	// debugging flags
	DEBUG_NET     = false
	DEBUG_CLIENT  = false
	DEBUG_CHANNEL = false
	DEBUG_SERVER  = false

	// errors
	ErrAlreadyDestroyed = errors.New("already destroyed")

	// regexps
	ChannelNameExpr = regexp.MustCompile(`^[&!#+][\pL\pN]{1,63}$`)
	NicknameExpr    = regexp.MustCompile(
		"^[\\pL\\[\\]{}^`][\\pL\\pN\\[\\]{}^`]{1,31}$")
)

const (
	SEM_VER       = "ergonomadic-1.2.11"
	CRLF          = "\r\n"
	MAX_REPLY_LEN = 512 - len(CRLF)

	LOGIN_TIMEOUT = time.Minute / 2 // how long the client has to login
	IDLE_TIMEOUT  = time.Minute     // how long before a client is considered idle
	QUIT_TIMEOUT  = time.Minute     // how long after idle before a client is kicked

	// string codes
	AWAY    StringCode = "AWAY"
	CAP     StringCode = "CAP"
	DEBUG   StringCode = "DEBUG"
	ERROR   StringCode = "ERROR"
	INVITE  StringCode = "INVITE"
	ISON    StringCode = "ISON"
	JOIN    StringCode = "JOIN"
	KICK    StringCode = "KICK"
	KILL    StringCode = "KILL"
	LIST    StringCode = "LIST"
	MODE    StringCode = "MODE"
	MOTD    StringCode = "MOTD"
	NAMES   StringCode = "NAMES"
	NICK    StringCode = "NICK"
	NOTICE  StringCode = "NOTICE"
	OPER    StringCode = "OPER"
	PART    StringCode = "PART"
	PASS    StringCode = "PASS"
	PING    StringCode = "PING"
	PONG    StringCode = "PONG"
	PRIVMSG StringCode = "PRIVMSG"
	PROXY   StringCode = "PROXY"
	QUIT    StringCode = "QUIT"
	TIME    StringCode = "TIME"
	TOPIC   StringCode = "TOPIC"
	USER    StringCode = "USER"
	VERSION StringCode = "VERSION"
	WHO     StringCode = "WHO"
	WHOIS   StringCode = "WHOIS"

	// numeric codes
	RPL_WELCOME           NumericCode = 1
	RPL_YOURHOST          NumericCode = 2
	RPL_CREATED           NumericCode = 3
	RPL_MYINFO            NumericCode = 4
	RPL_BOUNCE            NumericCode = 5
	RPL_TRACELINK         NumericCode = 200
	RPL_TRACECONNECTING   NumericCode = 201
	RPL_TRACEHANDSHAKE    NumericCode = 202
	RPL_TRACEUNKNOWN      NumericCode = 203
	RPL_TRACEOPERATOR     NumericCode = 204
	RPL_TRACEUSER         NumericCode = 205
	RPL_TRACESERVER       NumericCode = 206
	RPL_TRACESERVICE      NumericCode = 207
	RPL_TRACENEWTYPE      NumericCode = 208
	RPL_TRACECLASS        NumericCode = 209
	RPL_TRACERECONNECT    NumericCode = 210
	RPL_STATSLINKINFO     NumericCode = 211
	RPL_STATSCOMMANDS     NumericCode = 212
	RPL_ENDOFSTATS        NumericCode = 219
	RPL_UMODEIS           NumericCode = 221
	RPL_SERVLIST          NumericCode = 234
	RPL_SERVLISTEND       NumericCode = 235
	RPL_STATSUPTIME       NumericCode = 242
	RPL_STATSOLINE        NumericCode = 243
	RPL_LUSERCLIENT       NumericCode = 251
	RPL_LUSEROP           NumericCode = 252
	RPL_LUSERUNKNOWN      NumericCode = 253
	RPL_LUSERCHANNELS     NumericCode = 254
	RPL_LUSERME           NumericCode = 255
	RPL_ADMINME           NumericCode = 256
	RPL_ADMINLOC1         NumericCode = 257
	RPL_ADMINLOC2         NumericCode = 258
	RPL_ADMINEMAIL        NumericCode = 259
	RPL_TRACELOG          NumericCode = 261
	RPL_TRACEEND          NumericCode = 262
	RPL_TRYAGAIN          NumericCode = 263
	RPL_AWAY              NumericCode = 301
	RPL_USERHOST          NumericCode = 302
	RPL_ISON              NumericCode = 303
	RPL_UNAWAY            NumericCode = 305
	RPL_NOWAWAY           NumericCode = 306
	RPL_WHOISUSER         NumericCode = 311
	RPL_WHOISSERVER       NumericCode = 312
	RPL_WHOISOPERATOR     NumericCode = 313
	RPL_WHOWASUSER        NumericCode = 314
	RPL_ENDOFWHO          NumericCode = 315
	RPL_WHOISIDLE         NumericCode = 317
	RPL_ENDOFWHOIS        NumericCode = 318
	RPL_WHOISCHANNELS     NumericCode = 319
	RPL_LIST              NumericCode = 322
	RPL_LISTEND           NumericCode = 323
	RPL_CHANNELMODEIS     NumericCode = 324
	RPL_UNIQOPIS          NumericCode = 325
	RPL_NOTOPIC           NumericCode = 331
	RPL_TOPIC             NumericCode = 332
	RPL_INVITING          NumericCode = 341
	RPL_SUMMONING         NumericCode = 342
	RPL_INVITELIST        NumericCode = 346
	RPL_ENDOFINVITELIST   NumericCode = 347
	RPL_EXCEPTLIST        NumericCode = 348
	RPL_ENDOFEXCEPTLIST   NumericCode = 349
	RPL_VERSION           NumericCode = 351
	RPL_WHOREPLY          NumericCode = 352
	RPL_NAMREPLY          NumericCode = 353
	RPL_LINKS             NumericCode = 364
	RPL_ENDOFLINKS        NumericCode = 365
	RPL_ENDOFNAMES        NumericCode = 366
	RPL_BANLIST           NumericCode = 367
	RPL_ENDOFBANLIST      NumericCode = 368
	RPL_ENDOFWHOWAS       NumericCode = 369
	RPL_INFO              NumericCode = 371
	RPL_MOTD              NumericCode = 372
	RPL_ENDOFINFO         NumericCode = 374
	RPL_MOTDSTART         NumericCode = 375
	RPL_ENDOFMOTD         NumericCode = 376
	RPL_YOUREOPER         NumericCode = 381
	RPL_REHASHING         NumericCode = 382
	RPL_YOURESERVICE      NumericCode = 383
	RPL_TIME              NumericCode = 391
	RPL_USERSSTART        NumericCode = 392
	RPL_USERS             NumericCode = 393
	RPL_ENDOFUSERS        NumericCode = 394
	RPL_NOUSERS           NumericCode = 395
	ERR_NOSUCHNICK        NumericCode = 401
	ERR_NOSUCHSERVER      NumericCode = 402
	ERR_NOSUCHCHANNEL     NumericCode = 403
	ERR_CANNOTSENDTOCHAN  NumericCode = 404
	ERR_TOOMANYCHANNELS   NumericCode = 405
	ERR_WASNOSUCHNICK     NumericCode = 406
	ERR_TOOMANYTARGETS    NumericCode = 407
	ERR_NOSUCHSERVICE     NumericCode = 408
	ERR_NOORIGIN          NumericCode = 409
	ERR_INVALIDCAPCMD     NumericCode = 410
	ERR_NORECIPIENT       NumericCode = 411
	ERR_NOTEXTTOSEND      NumericCode = 412
	ERR_NOTOPLEVEL        NumericCode = 413
	ERR_WILDTOPLEVEL      NumericCode = 414
	ERR_BADMASK           NumericCode = 415
	ERR_UNKNOWNCOMMAND    NumericCode = 421
	ERR_NOMOTD            NumericCode = 422
	ERR_NOADMININFO       NumericCode = 423
	ERR_FILEERROR         NumericCode = 424
	ERR_NONICKNAMEGIVEN   NumericCode = 431
	ERR_ERRONEUSNICKNAME  NumericCode = 432
	ERR_NICKNAMEINUSE     NumericCode = 433
	ERR_NICKCOLLISION     NumericCode = 436
	ERR_UNAVAILRESOURCE   NumericCode = 437
	ERR_USERNOTINCHANNEL  NumericCode = 441
	ERR_NOTONCHANNEL      NumericCode = 442
	ERR_USERONCHANNEL     NumericCode = 443
	ERR_NOLOGIN           NumericCode = 444
	ERR_SUMMONDISABLED    NumericCode = 445
	ERR_USERSDISABLED     NumericCode = 446
	ERR_NOTREGISTERED     NumericCode = 451
	ERR_NEEDMOREPARAMS    NumericCode = 461
	ERR_ALREADYREGISTRED  NumericCode = 462
	ERR_NOPERMFORHOST     NumericCode = 463
	ERR_PASSWDMISMATCH    NumericCode = 464
	ERR_YOUREBANNEDCREEP  NumericCode = 465
	ERR_YOUWILLBEBANNED   NumericCode = 466
	ERR_KEYSET            NumericCode = 467
	ERR_CHANNELISFULL     NumericCode = 471
	ERR_UNKNOWNMODE       NumericCode = 472
	ERR_INVITEONLYCHAN    NumericCode = 473
	ERR_BANNEDFROMCHAN    NumericCode = 474
	ERR_BADCHANNELKEY     NumericCode = 475
	ERR_BADCHANMASK       NumericCode = 476
	ERR_NOCHANMODES       NumericCode = 477
	ERR_BANLISTFULL       NumericCode = 478
	ERR_NOPRIVILEGES      NumericCode = 481
	ERR_CHANOPRIVSNEEDED  NumericCode = 482
	ERR_CANTKILLSERVER    NumericCode = 483
	ERR_RESTRICTED        NumericCode = 484
	ERR_UNIQOPPRIVSNEEDED NumericCode = 485
	ERR_NOOPERHOST        NumericCode = 491
	ERR_UMODEUNKNOWNFLAG  NumericCode = 501
	ERR_USERSDONTMATCH    NumericCode = 502

	CAP_LS    CapSubCommand = "LS"
	CAP_LIST  CapSubCommand = "LIST"
	CAP_REQ   CapSubCommand = "REQ"
	CAP_ACK   CapSubCommand = "ACK"
	CAP_NAK   CapSubCommand = "NAK"
	CAP_CLEAR CapSubCommand = "CLEAR"
	CAP_END   CapSubCommand = "END"

	Add    ModeOp = '+'
	List   ModeOp = '='
	Remove ModeOp = '-'

	Away          UserMode = 'a'
	Invisible     UserMode = 'i'
	LocalOperator UserMode = 'O'
	Operator      UserMode = 'o'
	Restricted    UserMode = 'r'
	ServerNotice  UserMode = 's' // deprecated
	WallOps       UserMode = 'w'

	Anonymous       ChannelMode = 'a' // flag
	BanMask         ChannelMode = 'b' // arg
	ChannelCreator  ChannelMode = 'O' // flag
	ChannelOperator ChannelMode = 'o' // arg
	ExceptMask      ChannelMode = 'e' // arg
	InviteMask      ChannelMode = 'I' // arg
	InviteOnly      ChannelMode = 'i' // flag
	Key             ChannelMode = 'k' // flag arg
	Moderated       ChannelMode = 'm' // flag
	NoOutside       ChannelMode = 'n' // flag
	OpOnlyTopic     ChannelMode = 't' // flag
	Persistent      ChannelMode = 'P' // flag
	Private         ChannelMode = 'p' // flag
	Quiet           ChannelMode = 'q' // flag
	ReOp            ChannelMode = 'r' // flag
	Secret          ChannelMode = 's' // flag, deprecated
	UserLimit       ChannelMode = 'l' // flag arg
	Voice           ChannelMode = 'v' // arg

	MultiPrefix Capability = "multi-prefix"
	SASL        Capability = "sasl"

	Disable CapModifier = '-'
	Ack     CapModifier = '~'
	Sticky  CapModifier = '='
)

var (
	SupportedCapabilities = CapabilitySet{
		MultiPrefix: true,
	}
)

const (
	Registration Phase = iota
	Normal       Phase = iota
)

const (
	CapNone        CapState = iota
	CapNegotiating CapState = iota
	CapNegotiated  CapState = iota
)
