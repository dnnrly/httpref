package httpref

// RegisteredPorts is the list of all known IANA registered ports
var RegisteredPorts = References{
	{
		Name: "1024",
		Description: `
Reserved

IANA Status - Official
TCP         - Reserved
UDP         - Reserved

`,
	},
	{
		Name: "1027",
		Description: `
Reserved

IANA Status - Official
TCP         - Reserved
UDP         - Unspecified


"Native IPv6 behind IPv4-to-IPv4 NAT Customer Premises Equipment (6a44)"

IANA Status - Official
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "1028",
		Description: `
Deprecated

IANA Status - Official
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "1029",
		Description: `
"Microsoft DCOM (https://en.wikipedia.org/wiki/Distributed_Component_Object_Model) services"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "1058",
		Description: `
"nim, IBM (https://en.wikipedia.org/wiki/IBM) AIX (https://en.wikipedia.org/wiki/IBM_AIX) Network Installation Manager (https://en.wikipedia.org/wiki/Network_Installation_Manager) (NIM)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1059",
		Description: `
nimreg, IBM AIX Network Installation Manager (NIM)

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1080",
		Description: `
"SOCKS (https://en.wikipedia.org/wiki/SOCKS) proxy"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1085",
		Description: `
"WebObjects (https://en.wikipedia.org/wiki/WebObjects)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1098",
		Description: `
"rmiactivation, Java remote method invocation (https://en.wikipedia.org/wiki/Java_remote_method_invocation) (RMI) activation"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1099",
		Description: `
rmiregistry, Java remote method invocation (RMI) registry

IANA Status - Official
TCP         - Yes
UDP         - Assigned

`,
	},
	{
		Name: "1109",
		Description: `
Reserved – IANA

IANA Status - Official
TCP         - Unspecified
UDP         - Unspecified


"Kerberos (https://en.wikipedia.org/wiki/Kerberos_(protocol)) Post Office Protocol (KPOP (https://en.wikipedia.org/wiki/Kerberized_Post_Office_Protocol))"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "1113",
		Description: `
"Licklider Transmission Protocol (https://en.wikipedia.org/wiki/Licklider_Transmission_Protocol) (LTP) delay tolerant networking protocol"

IANA Status - Official
TCP         - "Assigned"
UDP         - "Yes"

`,
	},
	{
		Name: "1119",
		Description: `
"Battle.net (https://en.wikipedia.org/wiki/Battle.net) chat/game protocol, used by Blizzard (https://en.wikipedia.org/wiki/Blizzard_Entertainment)'s games"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1167",
		Description: `
"Cisco IP SLA (https://en.wikipedia.org/wiki/IP_SLA) (Service Assurance Agent)"

IANA Status - Official
TCP         - Yes, and SCTP
UDP         - Yes

`,
	},
	{
		Name: "1194",
		Description: `
"OpenVPN (https://en.wikipedia.org/wiki/OpenVPN#Networking)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1198",
		Description: `
"The cajo project (https://en.wikipedia.org/wiki/Cajo_project) Free dynamic transparent distributed computing in Java"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1214",
		Description: `
"Kazaa (https://en.wikipedia.org/wiki/Kazaa)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1218",
		Description: `
William POWER

IANA Status - Official
TCP         - Yes
UDP         - No

`,
	},
	{
		Name: "1220",
		Description: `
"QuickTime Streaming Server (https://en.wikipedia.org/wiki/QuickTime_Streaming_Server) administration"

IANA Status - Official
TCP         - Yes
UDP         - Assigned

`,
	},
	{
		Name: "1234",
		Description: `
"Infoseek (https://en.wikipedia.org/wiki/Infoseek) search agent"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"VLC media player (https://en.wikipedia.org/wiki/VLC_media_player) default port for UDP/RTP stream"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "1241",
		Description: `
"Nessus (https://en.wikipedia.org/wiki/Nessus_(software)) Security Scanner"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1270",
		Description: `
"Microsoft System Center Operations Manager (https://en.wikipedia.org/wiki/System_Center_Operations_Manager) (SCOM) (formerly Microsoft Operations Manager (MOM)) agent"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1293",
		Description: `
"Internet Protocol Security (IPSec (https://en.wikipedia.org/wiki/IPSec))"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1311",
		Description: `
Windows RxMon.exe

IANA Status - Official
TCP         - Yes
UDP         - Yes


"Dell OpenManage (https://en.wikipedia.org/wiki/OpenManage) HTTPS"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "1314",
		Description: `
"Festival Speech Synthesis System (https://en.wikipedia.org/wiki/Festival_Speech_Synthesis_System) server"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "1337",
		Description: `
"neo4j-shell (https://en.wikipedia.org/wiki/Neo4j)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Strapi (/w/index.php?title=Strapi&action=edit&redlink=1)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Sails.js (https://en.wikipedia.org/wiki/Sails.js) default port"

IANA Status - Unofficial
TCP         - Yes
UDP         - ?


"WASTE (https://en.wikipedia.org/wiki/WASTE) Encrypted File Sharing Program"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "1341",
		Description: `
"Qubes (Manufacturing Execution System (https://en.wikipedia.org/wiki/Manufacturing_Execution_System))"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1344",
		Description: `
"Internet Content Adaptation Protocol (https://en.wikipedia.org/wiki/Internet_Content_Adaptation_Protocol)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1352",
		Description: `
"IBM Lotus Notes (https://en.wikipedia.org/wiki/Lotus_Notes)/Domino (https://en.wikipedia.org/wiki/IBM_Lotus_Domino) (RPC) (https://en.wikipedia.org/wiki/Remote_procedure_call) protocol"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1360",
		Description: `
"Mimer SQL (https://en.wikipedia.org/wiki/Mimer_SQL)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1414",
		Description: `
"IBM (https://en.wikipedia.org/wiki/IBM) WebSphere MQ (https://en.wikipedia.org/wiki/WebSphere_MQ) (formerly known as MQSeries (https://en.wikipedia.org/wiki/MQSeries))"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1417",
		Description: `
"Timbuktu (https://en.wikipedia.org/wiki/Timbuktu_(software)) Service 1 Port"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1418",
		Description: `
Timbuktu Service 2 Port

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1419",
		Description: `
Timbuktu Service 3 Port

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1420",
		Description: `
Timbuktu Service 4 Port

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1431",
		Description: `
"Reverse Gossip Transport Protocol (/w/index.php?title=Reverse_Gossip_Transport_Protocol&action=edit&redlink=1) (RGTP), used to access a General-purpose Reverse-Ordered Gossip Gathering System (GROGGS) bulletin board (https://en.wikipedia.org/wiki/Bulletin_board_system), such as that implemented on the Cambridge University (https://en.wikipedia.org/wiki/University_of_Cambridge)'s Phoenix system (https://en.wikipedia.org/wiki/Phoenix_(computer))"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "1433",
		Description: `
"Microsoft SQL Server (https://en.wikipedia.org/wiki/Microsoft_SQL_Server) database management system (https://en.wikipedia.org/wiki/Database_management_system) (MSSQL) server"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1434",
		Description: `
Microsoft SQL Server database management system (MSSQL) monitor

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1481",
		Description: `
AIRS data interchange.

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1492",
		Description: `
"Sid Meier's CivNet (https://en.wikipedia.org/wiki/Sid_Meier%27s_CivNet), a multiplayer remake of the original Sid Meier's Civilization game"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "1494",
		Description: `
"Citrix Independent Computing Architecture (https://en.wikipedia.org/wiki/Independent_Computing_Architecture) (ICA)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1500",
		Description: `
"IBM Tivoli Storage Manager (https://en.wikipedia.org/wiki/IBM_Tivoli_Storage_Manager) server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "1501",
		Description: `
"IBM Tivoli Storage Manager client scheduler"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "1503",
		Description: `
"Windows Live Messenger (https://en.wikipedia.org/wiki/Windows_Live_Messenger) (Whiteboard and Application Sharing)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1512",
		Description: `
"Microsoft's Windows Internet Name Service (https://en.wikipedia.org/wiki/Windows_Internet_Name_Service) (WINS)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1513",
		Description: `
"Garena (https://en.wikipedia.org/wiki/Garena) game client"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1521",
		Description: `
"nCUBE (https://en.wikipedia.org/wiki/NCUBE) License Manager"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"Oracle database (https://en.wikipedia.org/wiki/Oracle_database) default listener, in future releases official port 2483 (TCP/IP) and 2484 (TCP/IP with SSL)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "1524",
		Description: `
"ingreslock, ingres (https://en.wikipedia.org/wiki/Ingres_(database))"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1527",
		Description: `
"Oracle Net Services (https://en.wikipedia.org/wiki/Oracle_Net_Services), formerly known as SQL*Net"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"Apache Derby Network Server (https://en.wikipedia.org/wiki/Apache_Derby#Derby_Network_Server)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "1533",
		Description: `
"IBM Sametime (https://en.wikipedia.org/wiki/IBM_Sametime) Virtual Places Chat"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1534",
		Description: `
"Eclipse Target Communication Framework "

IANA Status - Unofficial
TCP         - No
UDP         - Yes

`,
	},
	{
		Name: "1540",
		Description: `
"1C:Enterprise (https://en.wikipedia.org/wiki/1C:Enterprise) server agent (ragent)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1541",
		Description: `
"1C:Enterprise master cluster manager (rmngr)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1542",
		Description: `
"1C:Enterprise configuration repository server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1545",
		Description: `
"1C:Enterprise cluster administration server (RAS)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1547",
		Description: `
"Laplink (https://en.wikipedia.org/wiki/Laplink)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1550",
		Description: `
"1C:Enterprise debug server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes


"Gadu-Gadu (https://en.wikipedia.org/wiki/Gadu-Gadu) (direct client-to-client)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "1560–1590",
		Description: `
"1C:Enterprise cluster working processes"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1581",
		Description: `
"MIL STD 2045-47001 VMF (https://en.wikipedia.org/wiki/Combat-net_radio)"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"IBM Tivoli Storage Manager (https://en.wikipedia.org/wiki/IBM_Tivoli_Storage_Manager) web client"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "1582–1583",
		Description: `
"IBM Tivoli Storage Manager server web interface"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "1583",
		Description: `
"Pervasive PSQL (https://en.wikipedia.org/wiki/Pervasive_PSQL)"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "1589",
		Description: `
"Cisco VLAN Query Protocol (VQP (https://en.wikipedia.org/wiki/VQP))"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1604",
		Description: `
"DarkComet (https://en.wikipedia.org/wiki/DarkComet) remote administration tool (RAT)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1626",
		Description: `
"iSketch (https://en.wikipedia.org/wiki/ISketch)"

IANA Status - Unofficial
TCP         - Yes
UDP         - ?

`,
	},
	{
		Name: "1627",
		Description: `
"iSketch"

IANA Status - Unofficial
TCP         - Yes
UDP         - ?

`,
	},
	{
		Name: "1628",
		Description: `
"LonTalk (https://en.wikipedia.org/wiki/LonTalk) normal"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1629",
		Description: `
LonTalk urgent

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1645",
		Description: `
"Early deployment of RADIUS (https://en.wikipedia.org/wiki/RADIUS) before RFC standardization was done using UDP port number 1645. Enabled for compatibility reasons by default on Cisco (https://en.wikipedia.org/wiki/Cisco) and Juniper Networks (https://en.wikipedia.org/wiki/Juniper_Networks) RADIUS servers. Official port is 1812. TCP port 1645 must not be used."

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "1646",
		Description: `
"Old radacct port, RADIUS accounting protocol. Enabled for compatibility reasons by default on Cisco (https://en.wikipedia.org/wiki/Cisco) and Juniper Networks (https://en.wikipedia.org/wiki/Juniper_Networks) RADIUS servers. Official port is 1813. TCP port 1646 must not be used."

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "1666",
		Description: `
"Perforce (https://en.wikipedia.org/wiki/Perforce)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "1677",
		Description: `
"Novell GroupWise (https://en.wikipedia.org/wiki/Novell_GroupWise) clients in client/server access mode"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1688",
		Description: `
"Microsoft Key Management Service (https://en.wikipedia.org/wiki/Key_Management_Service) (KMS) for Windows Activation"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "1701",
		Description: `
"Layer 2 Forwarding Protocol (https://en.wikipedia.org/wiki/Layer_2_Forwarding_Protocol) (L2F)"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"Layer 2 Tunneling Protocol (https://en.wikipedia.org/wiki/Layer_2_Tunneling_Protocol) (L2TP)"

IANA Status - Official
TCP         - Assigned
UDP         - Yes

`,
	},
	{
		Name: "1707",
		Description: `
"Windward Studios (https://en.wikipedia.org/wiki/Windward_Studios) games (vdmplay)"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"L2TP/IPsec, for establish an initial connection"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "1716",
		Description: `
"America's Army (https://en.wikipedia.org/wiki/America%27s_Army), a massively multiplayer online game (https://en.wikipedia.org/wiki/Massively_multiplayer_online_game) (MMO)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "1719",
		Description: `
"H.323 (https://en.wikipedia.org/wiki/H.323) registration and alternate communication"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1720",
		Description: `
"H.323 (https://en.wikipedia.org/wiki/H.323) call signaling"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1723",
		Description: `
"Point-to-Point Tunneling Protocol (https://en.wikipedia.org/wiki/Point-to-Point_Tunneling_Protocol) (PPTP)"

IANA Status - Official
TCP         - Yes
UDP         - Assigned

`,
	},
	{
		Name: "1755",
		Description: `
"Microsoft Media Services (https://en.wikipedia.org/wiki/Microsoft_Media_Services) (MMS, ms-streaming)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1761",
		Description: `
"Novell ZENworks (https://en.wikipedia.org/wiki/Novell_ZENworks)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1783",
		Description: `
"Decomissioned [sic (https://en.wikipedia.org/wiki/Sic)] Port 04/14/00, ms"

IANA Status - Official
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "1801",
		Description: `
"Microsoft Message Queuing (https://en.wikipedia.org/wiki/Microsoft_Message_Queuing)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1812",
		Description: `
"RADIUS (https://en.wikipedia.org/wiki/RADIUS) authentication protocol, radius"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1813",
		Description: `
"RADIUS (https://en.wikipedia.org/wiki/RADIUS) accounting protocol, radius-acct"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1863",
		Description: `
"Microsoft Notification Protocol (https://en.wikipedia.org/wiki/Microsoft_Notification_Protocol) (MSNP), used by the Microsoft Messenger service (https://en.wikipedia.org/wiki/Microsoft_Messenger_service) and a number of instant messaging Messenger clients (https://en.wikipedia.org/wiki/Microsoft_Messenger_service#Official_clients)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1880",
		Description: `
"Node-RED (https://en.wikipedia.org/wiki/Node-RED)"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "1883",
		Description: `
"MQTT (https://en.wikipedia.org/wiki/MQTT) (formerly MQ Telemetry Transport)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1900",
		Description: `
"Simple Service Discovery Protocol (https://en.wikipedia.org/wiki/Simple_Service_Discovery_Protocol) (SSDP), discovery of UPnP (https://en.wikipedia.org/wiki/Universal_Plug_and_Play) devices"

IANA Status - Official
TCP         - Assigned
UDP         - Yes

`,
	},
	{
		Name: "1935",
		Description: `
"Macromedia Flash (https://en.wikipedia.org/wiki/Macromedia_Flash) Communications Server MX (https://en.wikipedia.org/wiki/Macromedia_Studio_MX), the precursor to Adobe Flash Media Server (https://en.wikipedia.org/wiki/Adobe_Flash_Media_Server) before Macromedia (https://en.wikipedia.org/wiki/Macromedia)'s acquisition by Adobe (https://en.wikipedia.org/wiki/Adobe_Inc) on December 3, 2005"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"Real Time Messaging Protocol (https://en.wikipedia.org/wiki/Real_Time_Messaging_Protocol) (RTMP), primarily used in Adobe Flash (https://en.wikipedia.org/wiki/Adobe_Flash)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1967",
		Description: `
"Cisco IOS IP Service Level Agreements (IP SLAs (https://en.wikipedia.org/wiki/IP_SLA)) Control Protocol"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "1970",
		Description: `
"Netop Remote Control (https://en.wikipedia.org/wiki/Netop_Remote_Control)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1972",
		Description: `
"InterSystems Caché (https://en.wikipedia.org/wiki/InterSystems_Cach%C3%A9)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1984",
		Description: `
"Big Brother (https://en.wikipedia.org/wiki/Big_Brother_(software))"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "1985",
		Description: `
"Cisco Hot Standby Router Protocol (https://en.wikipedia.org/wiki/Hot_Standby_Router_Protocol) (HSRP)"

IANA Status - Official
TCP         - Assigned
UDP         - Yes

`,
	},
	{
		Name: "1998",
		Description: `
"Cisco (https://en.wikipedia.org/wiki/Cisco) X.25 over TCP (XOT (https://en.wikipedia.org/wiki/XOT)) service"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2000",
		Description: `
"Cisco Skinny Client Control Protocol (https://en.wikipedia.org/wiki/Skinny_Client_Control_Protocol) (SCCP)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2010",
		Description: `
"Artemis: Spaceship Bridge Simulator (https://en.wikipedia.org/wiki/Artemis:_Spaceship_Bridge_Simulator)"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "2033",
		Description: `
"Civilization IV (https://en.wikipedia.org/wiki/Civilization_IV) multiplayer"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2049",
		Description: `
"Network File System (https://en.wikipedia.org/wiki/Network_File_System) (NFS)"

IANA Status - Official
TCP         - Yes, and SCTP
UDP         - Yes

`,
	},
	{
		Name: "2056",
		Description: `
"Civilization IV multiplayer"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2080",
		Description: `
"Autodesk (https://en.wikipedia.org/wiki/Autodesk) NLM (FLEXlm (https://en.wikipedia.org/wiki/FLEXlm))"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2082",
		Description: `
"cPanel (https://en.wikipedia.org/wiki/CPanel) default"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2083",
		Description: `
"Secure RADIUS (https://en.wikipedia.org/wiki/RADIUS) Service (radsec)"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"cPanel default SSL (https://en.wikipedia.org/wiki/Secure_Sockets_Layer)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2086",
		Description: `
"GNUnet (https://en.wikipedia.org/wiki/GNUnet)"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"WebHost Manager (https://en.wikipedia.org/wiki/WHM) default"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2087",
		Description: `
"WebHost Manager default SSL (https://en.wikipedia.org/wiki/Secure_Sockets_Layer)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2095",
		Description: `
"cPanel default web mail"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2096",
		Description: `
"cPanel default SSL (https://en.wikipedia.org/wiki/Secure_Sockets_Layer) web mail"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2100",
		Description: `
"Warzone 2100 (https://en.wikipedia.org/wiki/Warzone_2100) multiplayer"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2101",
		Description: `
"Networked Transport of RTCM via Internet Protocol (https://en.wikipedia.org/wiki/Networked_Transport_of_RTCM_via_Internet_Protocol) (NTRIP)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2102",
		Description: `
"Zephyr Notification Service (https://en.wikipedia.org/wiki/Zephyr_(protocol)) server"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2103",
		Description: `
Zephyr Notification Service serv-hm connection

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2104",
		Description: `
Zephyr Notification Service hostmanager

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2123",
		Description: `
"GTP (https://en.wikipedia.org/wiki/GPRS_Tunnelling_Protocol) control messages (GTP-C)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2142",
		Description: `
"TDMoIP (https://en.wikipedia.org/wiki/TDMoIP) (TDM over IP)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2152",
		Description: `
"GTP (https://en.wikipedia.org/wiki/GPRS_Tunnelling_Protocol) user data messages (GTP-U)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2159",
		Description: `
"GDB remote debug port (https://en.wikipedia.org/wiki/Gdbserver)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2181",
		Description: `
"EForward (https://en.wikipedia.org/wiki/EForward)-document transport system"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"Apache ZooKeeper (https://en.wikipedia.org/wiki/Apache_ZooKeeper) default client port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2195",
		Description: `
"Apple Push Notification Service (https://en.wikipedia.org/wiki/Apple_Push_Notification_Service)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2196",
		Description: `
"Apple Push Notification Service, feedback service"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2210",
		Description: `
"NOAAPORT (https://en.wikipedia.org/wiki/National_Weather_Service) Broadcast Network"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2211",
		Description: `
"EMWIN (https://en.wikipedia.org/wiki/EMWIN)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2221",
		Description: `
"ESET (https://en.wikipedia.org/wiki/ESET) anti-virus updates"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2222",
		Description: `
"EtherNet/IP (https://en.wikipedia.org/wiki/EtherNet/IP) implicit messaging for IO data"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"DirectAdmin (https://en.wikipedia.org/wiki/DirectAdmin) Access"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "2222–2226",
		Description: `
"ESET Remote administrator"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2240",
		Description: `
"General Dynamics (https://en.wikipedia.org/wiki/General_Dynamics) Remote Encryptor Configuration Information Protocol (RECIPe)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2261",
		Description: `
"CoMotion (https://en.wikipedia.org/wiki/CoMotion) master"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2262",
		Description: `
CoMotion backup

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2302",
		Description: `
"ArmA (https://en.wikipedia.org/wiki/ArmA) multiplayer"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes


"Halo: Combat Evolved (https://en.wikipedia.org/wiki/Halo:_Combat_Evolved) multiplayer host"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "2303",
		Description: `
"ArmA multiplayer (default port for game +1)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes


"Halo: Combat Evolved multiplayer listener"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "2305",
		Description: `
"ArmA multiplayer (default port for game +3)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "2351",
		Description: `
"AIM (https://en.wikipedia.org/wiki/AOL_Instant_Messenger) game LAN network port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2368",
		Description: `
"Ghost (blogging platform) (https://en.wikipedia.org/wiki/Ghost_(blogging_platform))"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2369",
		Description: `
"Default for BMC Control-M/Server (https://en.wikipedia.org/wiki/BMC_Control-M) Configuration Agent"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2370",
		Description: `
Default for BMC Control-M/Server, to allow the Control-M/Enterprise Manager to connect to the Control-M/Server

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2372",
		Description: `
"Default for K9 Web Protection (https://en.wikipedia.org/wiki/K9_Web_Protection)/parental controls, content filtering agent"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2375",
		Description: `
"Docker (https://en.wikipedia.org/wiki/Docker_(software)) REST API (plain)"

IANA Status - Official
TCP         - Yes
UDP         - Reserved

`,
	},
	{
		Name: "2376",
		Description: `
Docker REST API (SSL)

IANA Status - Official
TCP         - Yes
UDP         - Reserved

`,
	},
	{
		Name: "2377",
		Description: `
"Docker Swarm cluster management communications"

IANA Status - Official
TCP         - Yes
UDP         - Reserved

`,
	},
	{
		Name: "2379",
		Description: `
"CoreOS etcd (https://en.wikipedia.org/wiki/Etcd) client communication"

IANA Status - Official
TCP         - Yes
UDP         - Reserved


"KGS Go Server (https://en.wikipedia.org/wiki/KGS_Go_Server)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2380",
		Description: `
CoreOS etcd server communication

IANA Status - Official
TCP         - Yes
UDP         - Reserved

`,
	},
	{
		Name: "2389",
		Description: `
OpenView Session Mgr

IANA Status - Official
TCP         - Assigned
UDP         - Assigned

`,
	},
	{
		Name: "2399",
		Description: `
"FileMaker (https://en.wikipedia.org/wiki/FileMaker) Data Access Layer (ODBC/JDBC)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2401",
		Description: `
"CVS (https://en.wikipedia.org/wiki/Concurrent_Versions_System) version control system password-based server"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2404",
		Description: `
"IEC 60870-5-104 (https://en.wikipedia.org/wiki/IEC_60870-5-104), used to send electric power telecontrol messages between two systems via directly connected data circuits (/w/index.php?title=Data_circuit&action=edit&redlink=1)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2424",
		Description: `
"OrientDB (https://en.wikipedia.org/wiki/OrientDB) database listening for binary client connections"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2427",
		Description: `
"Media Gateway Control Protocol (https://en.wikipedia.org/wiki/Media_Gateway_Control_Protocol) (MGCP) media gateway"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2447",
		Description: `
"ovwdb—OpenView (https://en.wikipedia.org/wiki/OpenView) Network Node Manager (https://en.wikipedia.org/wiki/Network_Node_Manager) (NNM) daemon"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2459",
		Description: `
"XRPL (https://en.wikipedia.org/wiki/Xrp)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2480",
		Description: `
"OrientDB (https://en.wikipedia.org/wiki/OrientDB) database listening for HTTP client connections"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2483",
		Description: `
"Oracle database (https://en.wikipedia.org/wiki/Oracle_database) listening for insecure client connections to the listener, replaces port 1521"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2484",
		Description: `
"Oracle database listening for SSL (https://en.wikipedia.org/wiki/Secure_Sockets_Layer) client connections to the listener"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2500",
		Description: `
"NetFS communication"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2501",
		Description: `
NetFS probe

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "2535",
		Description: `
"Multicast Address Dynamic Client Allocation Protocol (https://en.wikipedia.org/wiki/Multicast_Address_Dynamic_Client_Allocation_Protocol) (MADCAP). All standard messages are UDP datagrams."

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2541",
		Description: `
"LonTalk (https://en.wikipedia.org/wiki/LonTalk)/IP"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2546–2548",
		Description: `
"EVault (https://en.wikipedia.org/wiki/EVault) data protection services"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2593",
		Description: `
"Ultima Online (https://en.wikipedia.org/wiki/Ultima_Online) servers"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2598",
		Description: `
"Citrix Independent Computing Architecture (https://en.wikipedia.org/wiki/Independent_Computing_Architecture) (ICA) with Session Reliability; port 1494 without session reliability"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2599",
		Description: `
"Ultima Online servers"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2628",
		Description: `
"DICT (https://en.wikipedia.org/wiki/DICT) "

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2638",
		Description: `
"SQL Anywhere (https://en.wikipedia.org/wiki/SQL_Anywhere) database server"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2710",
		Description: `
"XBT Tracker (https://en.wikipedia.org/wiki/XBT_Tracker). UDP tracker extension is considered experimental."

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2727",
		Description: `
"Media Gateway Control Protocol (https://en.wikipedia.org/wiki/Media_Gateway_Control_Protocol) (MGCP) media gateway controller (call agent)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2775",
		Description: `
"Short Message Peer-to-Peer (https://en.wikipedia.org/wiki/Short_Message_Peer-to-Peer) (SMPP)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2809",
		Description: `
"corbaloc:iiop URL, per the CORBA (https://en.wikipedia.org/wiki/CORBA) 3.0.3 specification"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2811",
		Description: `
"gsi ftp, per the GridFTP (https://en.wikipedia.org/wiki/GridFTP) specification"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2827",
		Description: `
"I2P (https://en.wikipedia.org/wiki/I2P) BOB Bridge"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "2944",
		Description: `
"Megaco (https://en.wikipedia.org/wiki/Megaco) text H.248"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2945",
		Description: `
Megaco binary (ASN.1) H.248

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2947",
		Description: `
"gpsd (https://en.wikipedia.org/wiki/Gpsd), GPS daemon"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2948",
		Description: `
"WAP (https://en.wikipedia.org/wiki/Wireless_Application_Protocol) push Multimedia Messaging Service (https://en.wikipedia.org/wiki/Multimedia_Messaging_Service) (MMS)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2949",
		Description: `
WAP push secure (MMS)

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "2967",
		Description: `
"Symantec System Center (https://en.wikipedia.org/wiki/Symantec_AntiVirus) agent (SSC-AGENT)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3000",
		Description: `
"Cloud9 IDE (https://en.wikipedia.org/wiki/Cloud9_IDE) server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Ruby on Rails (https://en.wikipedia.org/wiki/Ruby_on_Rails) development default"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Meteor (https://en.wikipedia.org/wiki/Meteor_(web_framework)) development default‹See TfM› (https://en.wikipedia.org/wiki/Wikipedia:Templates_for_discussion/Log/2020_August_1#Template:Failed_verification)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Resilio Sync (https://en.wikipedia.org/wiki/Resilio_Sync), spun from BitTorrent Sync."

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes


"Distributed Interactive Simulation (https://en.wikipedia.org/wiki/Distributed_Interactive_Simulation) (DIS)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "3004",
		Description: `
"iSync (https://en.wikipedia.org/wiki/ISync)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "3020",
		Description: `
"Common Internet File System (https://en.wikipedia.org/wiki/Common_Internet_File_System) (CIFS). See also port 445 for Server Message Block (https://en.wikipedia.org/wiki/Server_Message_Block) (SMB), a dialect of CIFS."

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3050",
		Description: `
"gds-db (Interbase (https://en.wikipedia.org/wiki/Interbase)/Firebird (https://en.wikipedia.org/wiki/Firebird_(database_server)) databases)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3052",
		Description: `
"APC (https://en.wikipedia.org/wiki/APC_by_Schneider_Electric) PowerChute Network (https://en.wikipedia.org/wiki/PowerChute)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3074",
		Description: `
"Xbox LIVE and Games for Windows – Live (https://en.wikipedia.org/wiki/Games_for_Windows_%E2%80%93_Live)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3101",
		Description: `
"BlackBerry Enterprise Server (https://en.wikipedia.org/wiki/BlackBerry_Enterprise_Server) communication protocol"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "3128",
		Description: `
"Squid (https://en.wikipedia.org/wiki/Squid_(software)) caching web proxy"

IANA Status - Unofficial
TCP         - Yes
UDP         - ?

`,
	},
	{
		Name: "3225",
		Description: `
"Fibre Channel over IP (https://en.wikipedia.org/wiki/Fibre_Channel_over_IP) (FCIP)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3233",
		Description: `
"WhiskerControl (https://en.wikipedia.org/wiki/WhiskerControl) research control protocol"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3260",
		Description: `
"iSCSI (https://en.wikipedia.org/wiki/ISCSI)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3268",
		Description: `
"msft-gc, Microsoft Global Catalog (LDAP (https://en.wikipedia.org/wiki/LDAP) service which contains data from Active Directory (https://en.wikipedia.org/wiki/Active_Directory) forests)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3269",
		Description: `
"msft-gc-ssl, Microsoft Global Catalog over SSL (https://en.wikipedia.org/wiki/Secure_Sockets_Layer) (similar to port 3268, LDAP (https://en.wikipedia.org/wiki/LDAP) over SSL (https://en.wikipedia.org/wiki/Secure_Sockets_Layer))"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3283",
		Description: `
"Net Assistant, a predecessor to Apple Remote Desktop"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"Apple Remote Desktop (https://en.wikipedia.org/wiki/Apple_Remote_Desktop) 2.0 or later"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3290",
		Description: `
"Virtual Air Traffic Simulation (https://en.wikipedia.org/wiki/VATSIM) (VATSIM) network voice communication"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "3305",
		Description: `
"Odette File Transfer Protocol (https://en.wikipedia.org/wiki/OFTP) (OFTP)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3306",
		Description: `
"MySQL (https://en.wikipedia.org/wiki/MySQL) database system"

IANA Status - Official
TCP         - Yes
UDP         - Assigned

`,
	},
	{
		Name: "3323",
		Description: `
"DECE (https://en.wikipedia.org/wiki/DECE) GEODI Server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3332",
		Description: `
Thundercloud DataPath Overlay Control

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "3333",
		Description: `
"Eggdrop (https://en.wikipedia.org/wiki/Eggdrop), an IRC bot default port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Network Caller ID (https://en.wikipedia.org/wiki/Network_Caller_ID) server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"CruiseControl.rb (https://en.wikipedia.org/wiki/CruiseControl.rb)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "3351",
		Description: `
"Pervasive PSQL (https://en.wikipedia.org/wiki/Pervasive_PSQL)"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "3386",
		Description: `
"GTP' (https://en.wikipedia.org/wiki/GTP%27) 3GPP (https://en.wikipedia.org/wiki/3GPP) GSM (https://en.wikipedia.org/wiki/GSM)/UMTS (https://en.wikipedia.org/wiki/UMTS) CDR (https://en.wikipedia.org/wiki/Call_detail_record) logging protocol"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3389",
		Description: `
"Microsoft Terminal Server (RDP (https://en.wikipedia.org/wiki/Remote_Desktop_Protocol)) officially registered as Windows Based Terminal (WBT)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3396",
		Description: `
"Novell (https://en.wikipedia.org/wiki/Novell) NDPS Printer Agent"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3412",
		Description: `
xmlBlaster

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3423",
		Description: `
Xware xTrm Communication Protocol

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "3424",
		Description: `
Xware xTrm Communication Protocol over SSL

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "3455",
		Description: `
"Resource Reservation Protocol (https://en.wikipedia.org/wiki/Resource_Reservation_Protocol) (RSVP)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3478",
		Description: `
"STUN (https://en.wikipedia.org/wiki/STUN), a protocol for NAT traversal"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"TURN (https://en.wikipedia.org/wiki/Traversal_Using_Relay_NAT), a protocol for NAT traversal (extension to STUN)"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"STUN Behavior Discovery. See also port 5349."

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3479",
		Description: `
"PlayStation Network (https://en.wikipedia.org/wiki/PlayStation_Network)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3480",
		Description: `
"PlayStation Network"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3483",
		Description: `
"Slim Devices (https://en.wikipedia.org/wiki/Slim_Devices) discovery protocol"

IANA Status - Official
TCP         - Unspecified
UDP         - Yes


"Slim Devices (https://en.wikipedia.org/wiki/Slim_Devices) SlimProto protocol"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "3493",
		Description: `
"Network UPS Tools (https://en.wikipedia.org/wiki/Network_UPS_Tools) (NUT)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3516",
		Description: `
Smartcard Port

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3527",
		Description: `
"Microsoft Message Queuing (https://en.wikipedia.org/wiki/Microsoft_Message_Queuing)"

IANA Status - Official
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "3535",
		Description: `
"SMTP (https://en.wikipedia.org/wiki/SMTP) alternate"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "3544",
		Description: `
"Teredo tunneling (https://en.wikipedia.org/wiki/Teredo_tunneling)"

IANA Status - Official
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "3632",
		Description: `
"Distcc (https://en.wikipedia.org/wiki/Distcc), distributed compiler"

IANA Status - Official
TCP         - Yes
UDP         - Assigned

`,
	},
	{
		Name: "3645",
		Description: `
"Cyc (https://en.wikipedia.org/wiki/Cyc)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3659",
		Description: `
"Apple SASL (https://en.wikipedia.org/wiki/Simple_Authentication_and_Security_Layer), used by Mac OS X Server (https://en.wikipedia.org/wiki/Mac_OS_X_Server) Password Server"

IANA Status - Official
TCP         - Yes
UDP         - Yes


Battlefield 4

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "3667",
		Description: `
Information Exchange

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3689",
		Description: `
"Digital Audio Access Protocol (https://en.wikipedia.org/wiki/Digital_Audio_Access_Protocol) (DAAP), used by Apple's (https://en.wikipedia.org/wiki/Apple_Inc.) iTunes (https://en.wikipedia.org/wiki/ITunes) and AirPlay (https://en.wikipedia.org/wiki/AirPlay)"

IANA Status - Official
TCP         - Yes
UDP         - Assigned

`,
	},
	{
		Name: "3690",
		Description: `
"Subversion (SVN) (https://en.wikipedia.org/wiki/Subversion_(software)) version control system"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3702",
		Description: `
"Web Services Dynamic Discovery (https://en.wikipedia.org/wiki/Web_Services_Dynamic_Discovery) (WS-Discovery), used by various components of Windows Vista (https://en.wikipedia.org/wiki/Windows_Vista) and later"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3724",
		Description: `
"Some Blizzard (https://en.wikipedia.org/wiki/Blizzard_Entertainment) games"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"Club Penguin (https://en.wikipedia.org/wiki/Club_Penguin) Disney online game for kids"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "3725",
		Description: `
Netia NA-ER Port

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3749",
		Description: `
"CimTrak (https://www.cimcor.com/cimtrak/) registered port"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3768",
		Description: `
RBLcheckd server daemon

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3784",
		Description: `
"Bidirectional Forwarding Detection (BFD)for IPv4 and IPv6 (Single Hop) (RFC 5881 (https://tools.ietf.org/html/rfc5881))"

IANA Status - Official
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "3785",
		Description: `
"VoIP program used by Ventrilo (https://en.wikipedia.org/wiki/Ventrilo)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "3799",
		Description: `
"RADIUS (https://en.wikipedia.org/wiki/RADIUS) change of authorization"

IANA Status - Official
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "3804",
		Description: `
"Harman Professional HiQnet (/w/index.php?title=HiQnet&action=edit&redlink=1) protocol"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3825",
		Description: `
"RedSeal Networks client/server connection"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "3826",
		Description: `
WarMUX game server

IANA Status - Official
TCP         - Yes
UDP         - Yes


"RedSeal Networks client/server connection"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "3835",
		Description: `
"RedSeal Networks client/server connection"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "3830",
		Description: `
System Management Agent, developed and used by Cerner to monitor and manage solutions

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3856",
		Description: `
ERP Server Application used by F10 Software

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3880",
		Description: `
IGRS

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3868",
		Description: `
"Diameter (https://en.wikipedia.org/wiki/Diameter_(protocol) ) base protocol (RFC 3588 (https://tools.ietf.org/html/rfc3588))"

IANA Status - Official
TCP         - Yes, and SCTP
UDP         - Unspecified

`,
	},
	{
		Name: "3872",
		Description: `
"Oracle Enterprise Manager (https://en.wikipedia.org/wiki/Oracle_Enterprise_Manager) Remote Agent"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "3900",
		Description: `
"udt_os, IBM UniData (https://en.wikipedia.org/wiki/IBM_U2) UDT OS"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "3960",
		Description: `
"Warframe (https://en.wikipedia.org/wiki/Warframe) online interaction"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "3962",
		Description: `
"Warframe online interaction"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "3978",
		Description: `
"OpenTTD (https://en.wikipedia.org/wiki/OpenTTD) game (masterserver and content service)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3979",
		Description: `
OpenTTD game

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "3999",
		Description: `
Norman distributed scanning service

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4000",
		Description: `
"Diablo II (https://en.wikipedia.org/wiki/Diablo_II) game"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4001",
		Description: `
"Microsoft Ants (https://en.wikipedia.org/wiki/Microsoft_Ants) game"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"CoreOS etcd (https://en.wikipedia.org/wiki/Etcd) client communication"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4018",
		Description: `
"Protocol information and warnings"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4035",
		Description: `
"IBM (https://en.wikipedia.org/wiki/IBM) Rational Developer for System z Remote System Explorer Daemon"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4045",
		Description: `
"Solaris (https://en.wikipedia.org/wiki/Solaris_(operating_system)) lockd NFS lock daemon/manager"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4050",
		Description: `
"Mud Master Chat protocol (MMCP) - Peer-to-peer communications between MUD (https://en.wikipedia.org/wiki/MUD) clients."

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4069",
		Description: `
"Minger Email Address Verification Protocol (https://en.wikipedia.org/wiki/Minger_Email_Address_Verification_Protocol)"

IANA Status - Official
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "4070",
		Description: `
"Amazon Echo (https://en.wikipedia.org/wiki/Amazon_Echo) Dot (Amazon Alexa (https://en.wikipedia.org/wiki/Amazon_Alexa)) streaming connection with Spotify (https://en.wikipedia.org/wiki/Spotify)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4089",
		Description: `
OpenCORE Remote Control Service

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4090",
		Description: `
"Kerio (https://en.wikipedia.org/wiki/Kerio_Technologies)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4093",
		Description: `
"PxPlus Client server interface ProvideX (https://en.wikipedia.org/wiki/ProvideX)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4096",
		Description: `
"Ascom Timeplex (https://en.wikipedia.org/wiki/Ascom_(company)) Bridge Relay Element (BRE)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4105",
		Description: `
Shofar (ShofarNexus)

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4111",
		Description: `
"Xgrid (https://en.wikipedia.org/wiki/Xgrid)"

IANA Status - Official
TCP         - Yes
UDP         - Assigned

`,
	},
	{
		Name: "4116",
		Description: `
Smartcard-TLS

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4125",
		Description: `
"Microsoft Remote Web Workplace (https://en.wikipedia.org/wiki/Microsoft_Remote_Web_Workplace) administration"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4172",
		Description: `
"Teradici (https://en.wikipedia.org/wiki/Teradici) PCoIP (https://en.wikipedia.org/wiki/PCoIP)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4190",
		Description: `
"ManageSieve"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4198",
		Description: `
"Couch Potato Android app"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4201",
		Description: `
"TinyMUD (https://en.wikipedia.org/wiki/TinyMUD) and various derivatives"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4222",
		Description: `
"NATS server default port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4226",
		Description: `
"Aleph One (https://en.wikipedia.org/wiki/Aleph_One_(computer_game)), a computer game"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4242",
		Description: `
"Orthanc (https://en.wikipedia.org/wiki/Orthanc_(software)) – DICOM (https://en.wikipedia.org/wiki/DICOM) server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Quassel (https://en.wikipedia.org/wiki/Quassel) distributed IRC client"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4243",
		Description: `
"Docker (https://en.wikipedia.org/wiki/Docker_(software)) implementations, redistributions, and setups default"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"CrashPlan (https://en.wikipedia.org/wiki/CrashPlan)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4244",
		Description: `
"Viber (https://en.wikipedia.org/wiki/Viber)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4303",
		Description: `
Simple Railroad Command Protocol (SRCP)

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4307",
		Description: `
"TrueConf (https://en.wikipedia.org/wiki/TrueConf) Client - TrueConf Server media data exchange"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4321",
		Description: `
"Referral Whois (RWhois) Protocol (https://en.wikipedia.org/wiki/RWhois)"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4444",
		Description: `
"Oracle (https://en.wikipedia.org/wiki/Oracle_Corporation) WebCenter Content: Content Server—Intradoc Socket port. (formerly known as Oracle Universal Content Management (https://en.wikipedia.org/wiki/Universal_Content_Management))."

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes


"Metasploit (https://en.wikipedia.org/wiki/Metasploit)'s default listener port"

IANA Status - Unofficial
TCP         - ?
UDP         - ?


"Xvfb (https://en.wikipedia.org/wiki/Xvfb) X server virtual frame buffer service"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4444–4445",
		Description: `
"I2P (https://en.wikipedia.org/wiki/I2P) HTTP/S proxy"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4486",
		Description: `
Integrated Client Message Service (ICMS)

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4488",
		Description: `
"Apple Wide Area Connectivity Service, used by Back to My Mac (https://en.wikipedia.org/wiki/Back_to_My_Mac)"

IANA Status - Official
TCP         - Yes
UDP         - Assigned

`,
	},
	{
		Name: "4500",
		Description: `
"IPSec NAT Traversal (https://en.wikipedia.org/wiki/IPsec_Passthrough) (RFC 3947 (https://tools.ietf.org/html/rfc3947), RFC 4306 (https://tools.ietf.org/html/rfc4306))"

IANA Status - Official
TCP         - Assigned
UDP         - Yes

`,
	},
	{
		Name: "4502–4534",
		Description: `
Microsoft Silverlight connectable ports under non-elevated trust

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4505–4506",
		Description: `
"Salt (https://en.wikipedia.org/wiki/Salt_(software)) master"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4534",
		Description: `
"Armagetron Advanced (https://en.wikipedia.org/wiki/Armagetron_Advanced) server default"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "4560",
		Description: `
"default Log4j (https://en.wikipedia.org/wiki/Log4j) socketappender port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4567",
		Description: `
"Sinatra (https://en.wikipedia.org/wiki/Sinatra_(software)) default server port in development mode (HTTP)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4569",
		Description: `
"Inter-Asterisk eXchange (https://en.wikipedia.org/wiki/Inter-Asterisk_eXchange) (IAX2)"

IANA Status - Official
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "4604",
		Description: `
"Identity Registration Protocol (https://en.wikipedia.org/wiki/Identity_Registration_Protocol)"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4605",
		Description: `
"Direct End to End Secure Chat Protocol (https://en.wikipedia.org/wiki/Direct_End_to_End_Secure_Chat_Protocol)"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4610–4640",
		Description: `
"QualiSystems (https://en.wikipedia.org/wiki/QualiSystems) TestShell Suite Services"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4662",
		Description: `
OrbitNet Message Service

IANA Status - Official
TCP         - Yes
UDP         - Yes


"Default for older versions of eMule (https://en.wikipedia.org/wiki/EMule)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4664",
		Description: `
"Google Desktop Search (https://en.wikipedia.org/wiki/Google_Desktop)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4672",
		Description: `
"Default for older versions of eMule (https://en.wikipedia.org/wiki/EMule)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "4711",
		Description: `
"eMule (https://en.wikipedia.org/wiki/EMule) optional web interface"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4713",
		Description: `
"PulseAudio (https://en.wikipedia.org/wiki/PulseAudio) sound server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4728",
		Description: `
"Computer Associates Desktop and Server Management (DMP)/Port Multiplexer"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4730",
		Description: `
"Gearman (https://en.wikipedia.org/wiki/Gearman)'s job server"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4739",
		Description: `
"IP Flow Information Export (https://en.wikipedia.org/wiki/IP_Flow_Information_Export)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4747",
		Description: `
"Apprentice (https://en.wikipedia.org/wiki/Apprentice_(software))"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4753",
		Description: `
SIMON (service and discovery)

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4789",
		Description: `
"Virtual eXtensible Local Area Network (VXLAN (https://en.wikipedia.org/wiki/Virtual_Extensible_LAN))"

IANA Status - Official
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "4791",
		Description: `
"IP Routable RocE (https://en.wikipedia.org/wiki/RDMA_over_Converged_Ethernet) (RoCEv2)"

IANA Status - Official
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "4840",
		Description: `
"OPC UA Connection Protocol (TCP) and OPC UA Multicast Datagram Protocol (UDP) for OPC Unified Architecture (https://en.wikipedia.org/wiki/OPC_Unified_Architecture) from OPC Foundation (https://en.wikipedia.org/wiki/OPC_Foundation)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4843",
		Description: `
"OPC UA TCP Protocol over TLS/SSL for OPC Unified Architecture (https://en.wikipedia.org/wiki/OPC_Unified_Architecture) from OPC Foundation (https://en.wikipedia.org/wiki/OPC_Foundation)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4847",
		Description: `
Web Fresh Communication, Quadrion Software & Odorless Entertainment

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4848",
		Description: `
Java, Glassfish Application Server administration default

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4894",
		Description: `
"LysKOM (https://en.wikipedia.org/wiki/LysKOM) Protocol A"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "4944",
		Description: `
"DrayTek (https://en.wikipedia.org/wiki/DrayTek) DSL Status Monitoring"

IANA Status - Unofficial
TCP         - No
UDP         - Yes

`,
	},
	{
		Name: "4949",
		Description: `
Munin Resource Monitoring Tool

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "4950",
		Description: `
Cylon Controls UC32 Communications Port

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5000",
		Description: `
"UPnP (https://en.wikipedia.org/wiki/Universal_Plug_and_Play)—Windows network device interoperability"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"VTun (https://en.wikipedia.org/wiki/VTun), VPN (https://en.wikipedia.org/wiki/VPN) Software"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes


"ASP.NET_Core (https://en.wikipedia.org/wiki/ASP.NET_Core) — Development Webserver"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"FlightGear (https://en.wikipedia.org/wiki/FlightGear) multiplayer"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified


"Synology Inc. (https://en.wikipedia.org/wiki/Synology_Inc.) Management Console, File Station, Audio Station"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Flask (https://en.wikipedia.org/wiki/Flask_(web_framework)) Development Webserver"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Heroku (https://en.wikipedia.org/wiki/Heroku) console access"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Docker (https://en.wikipedia.org/wiki/Docker_(software)) Registry"

IANA Status - Unofficial
TCP         - ?
UDP         - ?


"AT&T U-verse (https://en.wikipedia.org/wiki/AT%26T_U-verse) public, educational, and government access (https://en.wikipedia.org/wiki/Public,_educational,_and_government_access) (PEG) streaming over HTTP (https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"High-Speed SECS Message Services (https://en.wikipedia.org/wiki/High-Speed_SECS_Message_Services)"

IANA Status - Unofficial
TCP         - ?
UDP         - ?


"3CX Phone System (https://en.wikipedia.org/wiki/3CX_Phone_System) Legacy Management Console"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5000–5500",
		Description: `
"League of Legends (https://en.wikipedia.org/wiki/League_of_Legends), a multiplayer online battle arena (https://en.wikipedia.org/wiki/Multiplayer_online_battle_arena) video game"

IANA Status - Unofficial
TCP         - No
UDP         - Yes

`,
	},
	{
		Name: "5001",
		Description: `
"Slingbox (https://en.wikipedia.org/wiki/Slingbox) and Slingplayer"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Iperf (https://en.wikipedia.org/wiki/Iperf) (Tool for measuring TCP and UDP bandwidth performance)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes


"Synology Inc. (https://en.wikipedia.org/wiki/Synology_Inc.) Secured Management Console, File Station, Audio Station"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"3CX Phone System (https://en.wikipedia.org/wiki/3CX_Phone_System) Secured Management Console, Secure API"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5002",
		Description: `
"ASSA ARX access control system"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5003",
		Description: `
"FileMaker (https://en.wikipedia.org/wiki/FileMaker) – name binding and transport"

IANA Status - Official
TCP         - Yes
UDP         - Assigned

`,
	},
	{
		Name: "5004",
		Description: `
"Real-time Transport Protocol (https://en.wikipedia.org/wiki/Real-time_Transport_Protocol) media data (RTP) (RFC 3551 (https://tools.ietf.org/html/rfc3551), RFC 4571 (https://tools.ietf.org/html/rfc4571))"

IANA Status - Official
TCP         - Yes, and DCCP
UDP         - Yes

`,
	},
	{
		Name: "5005",
		Description: `
"Real-time Transport Protocol control protocol (https://en.wikipedia.org/wiki/RTP_Control_Protocol) (RTCP) (RFC 3551 (https://tools.ietf.org/html/rfc3551), RFC 4571 (https://tools.ietf.org/html/rfc4571))"

IANA Status - Official
TCP         - Yes, and DCCP
UDP         - Yes

`,
	},
	{
		Name: "5007",
		Description: `
Palo Alto Networks - User-ID agent

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5010",
		Description: `
"Registered to: TelePath (the IBM FlowMark workflow-management system (https://en.wikipedia.org/wiki/Workflow-management_system) messaging platform)
The TCP port is now used for: IBM WebSphere MQ (https://en.wikipedia.org/wiki/WebSphere_MQ) Workflow"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5011",
		Description: `
"TelePath (the IBM FlowMark workflow-management system (https://en.wikipedia.org/wiki/Workflow-management_system) messaging platform)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5025",
		Description: `
"scpi-raw Standard Commands for Programmable Instruments (https://en.wikipedia.org/wiki/Standard_Commands_for_Programmable_Instruments)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5029",
		Description: `
Sonic Robo Blast 2 and Sonic Robo Blast 2 Kart servers

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "5031",
		Description: `
"AVM CAPI-over-TCP (ISDN (https://en.wikipedia.org/wiki/ISDN) over Ethernet (https://en.wikipedia.org/wiki/Ethernet) tunneling)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5037",
		Description: `
Android ADB server

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5044",
		Description: `
Standard port in Filebeats/Logstash implementation of Lumberjack protocol.

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5048",
		Description: `
Texai Message Service

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5050",
		Description: `
"Yahoo! Messenger (https://en.wikipedia.org/wiki/Yahoo!_Messenger)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5051",
		Description: `
"ita-agent Symantec (https://en.wikipedia.org/wiki/NortonLifeLock) Intruder Alert"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5060",
		Description: `
"Session Initiation Protocol (https://en.wikipedia.org/wiki/Session_Initiation_Protocol) (SIP)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5061",
		Description: `
"Session Initiation Protocol (https://en.wikipedia.org/wiki/Session_Initiation_Protocol) (SIP) over TLS (https://en.wikipedia.org/wiki/Transport_Layer_Security)"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5062",
		Description: `
Localisation access

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5064",
		Description: `
"EPICS (https://en.wikipedia.org/wiki/EPICS) Channel Access server"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5065",
		Description: `
"EPICS Channel Access repeater beacon"

IANA Status - Official
TCP         - Assigned
UDP         - Yes

`,
	},
	{
		Name: "5070",
		Description: `
"Binary Floor Control Protocol (/w/index.php?title=Binary_Floor_Control_Protocol&action=edit&redlink=1) (BFCP)"

IANA Status - Unofficial
TCP         - Yes
UDP         - No

`,
	},
	{
		Name: "5084",
		Description: `
"EPCglobal (https://en.wikipedia.org/wiki/EPCglobal) Low Level Reader Protocol (LLRP (https://en.wikipedia.org/wiki/LLRP))"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5085",
		Description: `
"EPCglobal Low Level Reader Protocol (LLRP (https://en.wikipedia.org/wiki/LLRP)) over TLS (https://en.wikipedia.org/wiki/Transport_Layer_Security)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5090",
		Description: `
"3CX Phone System (https://en.wikipedia.org/wiki/3CX_Phone_System) 3CX Tunnel Protocol, 3CX App API, 3CX Session Border Controller"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5093",
		Description: `
"SafeNet, Inc (https://en.wikipedia.org/wiki/SafeNet) Sentinel LM, Sentinel RMS, License Manager, client-to-server"

IANA Status - Official
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "5099",
		Description: `
SafeNet, Inc Sentinel LM, Sentinel RMS, License Manager, server-to-server

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5104",
		Description: `
"IBM (https://en.wikipedia.org/wiki/IBM) Tivoli Framework (/w/index.php?title=IBM_Tivoli_Framework&action=edit&redlink=1) NetCOOL/Impact HTTP (https://en.wikipedia.org/wiki/HTTP) Service"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5121",
		Description: `
"Neverwinter Nights (https://en.wikipedia.org/wiki/Neverwinter_Nights)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5124",
		Description: `
"TorgaNET (Micronational (https://en.wikipedia.org/wiki/Micronation) Darknet (https://en.wikipedia.org/wiki/Darknet))"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5125",
		Description: `
"TorgaNET (Micronational (https://en.wikipedia.org/wiki/Micronation) Intelligence Darknet (https://en.wikipedia.org/wiki/Darknet))"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5150",
		Description: `
"ATMP Ascend Tunnel Management Protocol"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5151",
		Description: `
"ESRI (https://en.wikipedia.org/wiki/Environmental_Systems_Research_Institute) SDE Instance"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified


ESRI SDE Remote Start

IANA Status - Official
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "5154",
		Description: `
"BZFlag (https://en.wikipedia.org/wiki/BZFlag)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5172",
		Description: `
"PC over IP Endpoint Management"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5190",
		Description: `
"AOL Instant Messenger (https://en.wikipedia.org/wiki/AOL_Instant_Messenger) protocol. The chat app is defunct as of 15 December 2017."

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5198",
		Description: `
"EchoLink (https://en.wikipedia.org/wiki/Echolink) VoIP Amateur Radio Software (Voice)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "5199",
		Description: `
EchoLink VoIP Amateur Radio Software (Voice)

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "5200",
		Description: `
EchoLink VoIP Amateur Radio Software (Information)

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5201",
		Description: `
"Iperf3 (https://en.wikipedia.org/wiki/Iperf3) (Tool for measuring TCP and UDP bandwidth performance)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5222",
		Description: `
"Extensible Messaging and Presence Protocol (https://en.wikipedia.org/wiki/Extensible_Messaging_and_Presence_Protocol) (XMPP) client connection"

IANA Status - Official
TCP         - Yes
UDP         - Reserved

`,
	},
	{
		Name: "5223",
		Description: `
"Apple Push Notification Service (https://en.wikipedia.org/wiki/Apple_Push_Notification_Service)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Extensible Messaging and Presence Protocol (XMPP) client connection over SSL (https://en.wikipedia.org/wiki/Secure_Sockets_Layer)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5228",
		Description: `
HP Virtual Room Service

IANA Status - Official
TCP         - Yes
UDP         - Unspecified


"Google Play (https://en.wikipedia.org/wiki/Google_Play), Android Cloud to Device Messaging Service (https://en.wikipedia.org/wiki/Android_Cloud_to_Device_Messaging_Service), Google Cloud Messaging (https://en.wikipedia.org/wiki/Google_Cloud_Messaging)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5242",
		Description: `
"Viber (https://en.wikipedia.org/wiki/Viber)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5243",
		Description: `
"Viber"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5246",
		Description: `
"Control And Provisioning of Wireless Access Points (CAPWAP (https://en.wikipedia.org/wiki/CAPWAP)) CAPWAP control"

IANA Status - Official
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "5247",
		Description: `
"Control And Provisioning of Wireless Access Points (CAPWAP) CAPWAP data"

IANA Status - Official
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "5269",
		Description: `
"Extensible Messaging and Presence Protocol (XMPP) server-to-server connection"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5280",
		Description: `
"Extensible Messaging and Presence Protocol (XMPP)"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5281",
		Description: `
"Extensible Messaging and Presence Protocol (XMPP)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5298",
		Description: `
"Extensible Messaging and Presence Protocol (XMPP)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5310",
		Description: `
"Outlaws (https://en.wikipedia.org/wiki/Outlaws_(1997_video_game)), a 1997 first-person shooter video game"

IANA Status - Official
TCP         - Assigned
UDP         - Yes

`,
	},
	{
		Name: "5318",
		Description: `
"Certificate Management over CMS (https://en.wikipedia.org/wiki/Certificate_Management_over_CMS)"

IANA Status - Official
TCP         - Yes
UDP         - Reserved

`,
	},
	{
		Name: "5349",
		Description: `
"STUN (https://en.wikipedia.org/wiki/STUN) over TLS (https://en.wikipedia.org/wiki/Transport_Layer_Security)/DTLS (https://en.wikipedia.org/wiki/Datagram_Transport_Layer_Security), a protocol for NAT traversal (https://en.wikipedia.org/wiki/NAT_traversal)"

IANA Status - Official
TCP         - Yes/No
UDP         - Yes/No


"TURN (https://en.wikipedia.org/wiki/Traversal_Using_Relay_NAT) over TLS/DTLS, a protocol for NAT traversal"

IANA Status - Official
TCP         - Yes/No
UDP         - Yes/No


"STUN Behavior Discovery over TLS. See also port 3478."

IANA Status - Official
TCP         - Yes
UDP         - Reserved

`,
	},
	{
		Name: "5351",
		Description: `
"NAT Port Mapping Protocol (https://en.wikipedia.org/wiki/NAT_Port_Mapping_Protocol) and Port Control Protocol (https://en.wikipedia.org/wiki/Port_Control_Protocol)—client-requested configuration for connections through network address translators (https://en.wikipedia.org/wiki/Network_Address_Translation) and firewalls (https://en.wikipedia.org/wiki/Firewall_(computing))"

IANA Status - Official
TCP         - Reserved
UDP         - Yes

`,
	},
	{
		Name: "5353",
		Description: `
"Multicast DNS (https://en.wikipedia.org/wiki/Multicast_DNS) (mDNS)"

IANA Status - Official
TCP         - Assigned
UDP         - Yes

`,
	},
	{
		Name: "5355",
		Description: `
"Link-Local Multicast Name Resolution (https://en.wikipedia.org/wiki/LLMNR) (LLMNR), allows hosts (https://en.wikipedia.org/wiki/Host_(network)) to perform name resolution (https://en.wikipedia.org/wiki/Hostname_resolution) for hosts on the same local link (https://en.wikipedia.org/wiki/Local_area_network) (only provided by Windows Vista (https://en.wikipedia.org/wiki/Windows_Vista) and Server 2008 (https://en.wikipedia.org/wiki/Server_2008))"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5357",
		Description: `
"Web Services for Devices (https://en.wikipedia.org/wiki/Web_Services_for_Devices) (WSDAPI) (only provided by Windows Vista, Windows 7 and Server 2008)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5358",
		Description: `
"WSDAPI (https://en.wikipedia.org/wiki/Web_Services_for_Devices) Applications to Use a Secure Channel (only provided by Windows Vista, Windows 7 and Server 2008)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5394",
		Description: `
"Kega Fusion, a Sega multi-console emulator"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "5402",
		Description: `
"Multicast File Transfer Protocol (/w/index.php?title=Multicast_File_Transfer_Protocol&action=edit&redlink=1) (MFTP)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5405",
		Description: `
"NetSupport Manager (https://en.wikipedia.org/wiki/NetSupport_Manager)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5412",
		Description: `
"IBM (https://en.wikipedia.org/wiki/IBM) Rational Synergy (Telelogic Synergy (https://en.wikipedia.org/wiki/Telelogic_Synergy)) (Continuus CM) Message Router"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5413",
		Description: `
"Wonderware (https://en.wikipedia.org/wiki/Wonderware) SuiteLink service"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5417",
		Description: `
SNS Agent

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5421",
		Description: `
"NetSupport Manager (https://en.wikipedia.org/wiki/NetSupport_Manager)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5432",
		Description: `
"PostgreSQL (https://en.wikipedia.org/wiki/PostgreSQL) database system"

IANA Status - Official
TCP         - Yes
UDP         - Assigned

`,
	},
	{
		Name: "5433",
		Description: `
"Bouwsoft file/webserver"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5445",
		Description: `
"Cisco Unified Video Advantage"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "5480",
		Description: `
"VMware (https://en.wikipedia.org/wiki/VMware) VAMI (Virtual Appliance Management Infrastructure)—used for initial setup of various administration settings on Virtual Appliances designed using the VAMI architecture."

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5481",
		Description: `
"Schneider Electric (https://en.wikipedia.org/wiki/Schneider_Electric)'s ClearSCADA (SCADA (https://en.wikipedia.org/wiki/SCADA) implementation for Windows) — used for client-to-server communication."

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5495",
		Description: `
"IBM Cognos TM1 (https://en.wikipedia.org/wiki/TM1) Admin server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5498",
		Description: `
"Hotline (https://en.wikipedia.org/wiki/Hotline_Communications) tracker server connection"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5499",
		Description: `
Hotline tracker server discovery

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "5500",
		Description: `
Hotline control connection

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"VNC (https://en.wikipedia.org/wiki/Virtual_Network_Computing) Remote Frame Buffer RFB protocol (https://en.wikipedia.org/wiki/RFB_protocol)—for incoming listening viewer"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5501",
		Description: `
Hotline file transfer connection

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5517",
		Description: `
"Setiqueue (https://en.wikipedia.org/wiki/SETI@home#Software) Proxy server client for SETI@Home (https://en.wikipedia.org/wiki/SETI@Home) project"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5550",
		Description: `
"Hewlett-Packard (https://en.wikipedia.org/wiki/Hewlett-Packard) Data Protector"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5554",
		Description: `
"Fastboot (https://en.wikipedia.org/wiki/Fastboot) default wireless port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5555",
		Description: `
"Oracle (https://en.wikipedia.org/wiki/Oracle_Corporation) WebCenter Content: Inbound Refinery—Intradoc Socket port. (formerly known as Oracle Universal Content Management (https://en.wikipedia.org/wiki/Universal_Content_Management)). Port though often changed during installation"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes


"Freeciv (https://en.wikipedia.org/wiki/Freeciv) versions up to 2.0, Hewlett-Packard (https://en.wikipedia.org/wiki/Hewlett-Packard) Data Protector, McAfee EndPoint Encryption (https://en.wikipedia.org/wiki/Comparison_of_disk_encryption_software) Database Server, SAP (https://en.wikipedia.org/wiki/Session_Announcement_Protocol), Default for Microsoft Dynamics CRM 4.0, Softether VPN default port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5556",
		Description: `
"Freeciv (https://en.wikipedia.org/wiki/Freeciv), Oracle WebLogic Server Node Manager"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5568",
		Description: `
"Session Data Transport (SDT), a part of Architecture for Control Networks (https://en.wikipedia.org/wiki/Architecture_for_Control_Networks) (ACN)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5601",
		Description: `
"Kibana (https://en.wikipedia.org/wiki/Kibana)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5631",
		Description: `
"pcANYWHEREdata, Symantec (https://en.wikipedia.org/wiki/NortonLifeLock) pcAnywhere (https://en.wikipedia.org/wiki/Pcanywhere) (version 7.52 and later) data"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5632",
		Description: `
pcANYWHEREstat, Symantec pcAnywhere (version 7.52 and later) status

IANA Status - Official
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "5656",
		Description: `
"IBM Lotus Sametime (https://en.wikipedia.org/wiki/IBM_Lotus_Sametime) p2p file transfer"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5666",
		Description: `
"NRPE (https://en.wikipedia.org/wiki/NRPE) (Nagios (https://en.wikipedia.org/wiki/Nagios))"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5667",
		Description: `
NSCA (Nagios)

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5670",
		Description: `
"FILEMQ (/w/index.php?title=FILEMQ&action=edit&redlink=1) ZeroMQ File Message Queuing Protocol"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified


"ZRE-DISC (/w/index.php?title=ZRE-DISC&action=edit&redlink=1) ZeroMQ Realtime Exchange Protocol (Discovery)"

IANA Status - Official
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "5671",
		Description: `
"Advanced Message Queuing Protocol (https://en.wikipedia.org/wiki/Advanced_Message_Queuing_Protocol) (AMQP) over TLS (https://en.wikipedia.org/wiki/Transport_Layer_Security)"

IANA Status - Official
TCP         - Yes
UDP         - Assigned

`,
	},
	{
		Name: "5672",
		Description: `
"Advanced Message Queuing Protocol (AMQP)"

IANA Status - Official
TCP         - Yes, and SCTP
UDP         - Assigned

`,
	},
	{
		Name: "5683",
		Description: `
"Constrained Application Protocol (https://en.wikipedia.org/wiki/Constrained_Application_Protocol) (CoAP)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5684",
		Description: `
Constrained Application Protocol Secure (CoAPs)

IANA Status - Official
TCP         - Yes/No
UDP         - Yes

`,
	},
	{
		Name: "5693",
		Description: `
"Nagios (https://en.wikipedia.org/wiki/Nagios) Cross Platform Agent (NCPA)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5701",
		Description: `
"Hazelcast (https://en.wikipedia.org/wiki/Hazelcast) default communication port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5722",
		Description: `
"Microsoft RPC, DFSR (SYSVOL) Replication Service"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5718",
		Description: `
Microsoft DPM Data Channel (with the agent coordinator)

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5719",
		Description: `
Microsoft DPM Data Channel (with the protection agent)

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5723",
		Description: `
"System Center Operations Manager (https://en.wikipedia.org/wiki/System_Center_Operations_Manager)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5724",
		Description: `
Operations Manager Console

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5741",
		Description: `
IDA Discover Port 1

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5742",
		Description: `
IDA Discover Port 2

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5800",
		Description: `
"VNC (https://en.wikipedia.org/wiki/Virtual_Network_Computing) Remote Frame Buffer RFB protocol (https://en.wikipedia.org/wiki/RFB_protocol) over HTTP (https://en.wikipedia.org/wiki/HTTP)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"ProjectWise Server (https://en.wikipedia.org/wiki/ProjectWise)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5900",
		Description: `
"Remote Frame Buffer protocol (https://en.wikipedia.org/wiki/RFB_protocol) (RFB)"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"Virtual Network Computing (https://en.wikipedia.org/wiki/Virtual_Network_Computing) (VNC) Remote Frame Buffer RFB protocol (https://en.wikipedia.org/wiki/RFB_protocol)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5931",
		Description: `
"AMMYY (https://en.wikipedia.org/wiki/AMMYY) admin Remote Control"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5938",
		Description: `
"TeamViewer (https://en.wikipedia.org/wiki/TeamViewer) remote desktop protocol"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5984",
		Description: `
"CouchDB (https://en.wikipedia.org/wiki/CouchDB) database server"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "5985",
		Description: `
"Windows PowerShell (https://en.wikipedia.org/wiki/Windows_PowerShell) Default psSession Port Windows Remote Management Service (https://en.wikipedia.org/wiki/WS-Management) (WinRM-HTTP)"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5986",
		Description: `
"Windows PowerShell Default psSession Port Windows Remote Management Service (https://en.wikipedia.org/wiki/WS-Management) (WinRM-HTTPS)"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "5988–5989",
		Description: `
"CIM (https://en.wikipedia.org/wiki/Common_Information_Model_(computing))-XML (DMTF Protocol)"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6000–6063",
		Description: `
"X11 (https://en.wikipedia.org/wiki/X_Window_System)—used between an X client and server over the network"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6005",
		Description: `
"Default for BMC Software (https://en.wikipedia.org/wiki/BMC_Software) Control-M/Server (https://en.wikipedia.org/wiki/BMC_Control-M)—Socket used for communication between Control-M processes—though often changed during installation"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Default for Camfrog (https://en.wikipedia.org/wiki/Camfrog) chat & cam client"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6009",
		Description: `
"JD Edwards EnterpriseOne (https://en.wikipedia.org/wiki/JD_Edwards_EnterpriseOne) ERP system JDENet messaging client listener"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6050",
		Description: `
"Arcserve (https://en.wikipedia.org/wiki/Arcserve) backup"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6051",
		Description: `
Arcserve backup

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6086",
		Description: `
"Peer Distributed Transfer Protocol (https://en.wikipedia.org/wiki/Peer_Distributed_Transfer_Protocol) (PDTP), FTP like file server in a P2P network"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6100",
		Description: `
"Vizrt (https://en.wikipedia.org/wiki/Vizrt) System"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Ventrilo (https://en.wikipedia.org/wiki/Ventrilo) authentication for version 3"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6101",
		Description: `
"Backup Exec Agent Browser"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6110",
		Description: `
"softcm, HP (https://en.wikipedia.org/wiki/Hewlett-Packard) Softbench (https://en.wikipedia.org/wiki/Softbench) CM"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6111",
		Description: `
"spc, HP (https://en.wikipedia.org/wiki/Hewlett-Packard) Softbench (https://en.wikipedia.org/wiki/Softbench) Sub-Process Control"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6112",
		Description: `
dtspcd, execute commands and launch applications remotely

IANA Status - Official
TCP         - Yes
UDP         - Yes


"Blizzard (https://en.wikipedia.org/wiki/Blizzard_Entertainment)'s Battle.net (https://en.wikipedia.org/wiki/Battle.net) gaming service and some games, ArenaNet (https://en.wikipedia.org/wiki/ArenaNet) gaming service, Relic (https://en.wikipedia.org/wiki/Relic_Entertainment) gaming service"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes


"Club Penguin (https://en.wikipedia.org/wiki/Club_Penguin) Disney online game for kids"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6113",
		Description: `
"Club Penguin (https://en.wikipedia.org/wiki/Club_Penguin) Disney online game for kids, Used by some Blizzard (https://en.wikipedia.org/wiki/Blizzard_Entertainment) games"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6136",
		Description: `
"ObjectDB (https://en.wikipedia.org/wiki/ObjectDB) database server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6159",
		Description: `
"ARINC (https://en.wikipedia.org/wiki/ARINC) 840 EFB (https://en.wikipedia.org/wiki/Electronic_flight_bag) Application Control Interface"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6200",
		Description: `
"Oracle WebCenter (https://en.wikipedia.org/wiki/Oracle_WebCenter) Content Portable: Content Server (With Native UI) and Inbound Refinery"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6201",
		Description: `
Oracle WebCenter Content Portable: Admin

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6225",
		Description: `
Oracle WebCenter Content Portable: Content Server Web UI

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6227",
		Description: `
Oracle WebCenter Content Portable: JavaDB

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6240",
		Description: `
Oracle WebCenter Content Portable: Capture

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6244",
		Description: `
Oracle WebCenter Content Portable: Content Server—Intradoc Socket port

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6255",
		Description: `
Oracle WebCenter Content Portable: Inbound Refinery—Intradoc Socket port

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6257",
		Description: `
"WinMX (https://en.wikipedia.org/wiki/WinMX) (see also 6699)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "6260",
		Description: `
planet M.U.L.E.

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6262",
		Description: `
"Sybase Advantage Database Server (https://en.wikipedia.org/wiki/Advantage_Database_Server)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6343",
		Description: `
"SFlow (https://en.wikipedia.org/wiki/SFlow), sFlow traffic monitoring"

IANA Status - Official
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "6346",
		Description: `
"gnutella-svc (https://en.wikipedia.org/wiki/Gnutella), gnutella (FrostWire (https://en.wikipedia.org/wiki/FrostWire), Limewire (https://en.wikipedia.org/wiki/Limewire), Shareaza (https://en.wikipedia.org/wiki/Shareaza), etc.)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6347",
		Description: `
gnutella-rtr, Gnutella alternate

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6350",
		Description: `
App Discovery and Access Protocol

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6379",
		Description: `
"Redis (https://en.wikipedia.org/wiki/Redis) key-value data store"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6389",
		Description: `
"EMC (https://en.wikipedia.org/wiki/EMC_Corporation) CLARiiON (https://en.wikipedia.org/wiki/CLARiiON)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6432",
		Description: `
PgBouncer—A connection pooler for PostgreSQL

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6436",
		Description: `
"Leap Motion (https://en.wikipedia.org/wiki/Leap_Motion) Websocket Server TLS"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6437",
		Description: `
Leap Motion Websocket Server

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6444",
		Description: `
"Sun Grid Engine (https://en.wikipedia.org/wiki/Sun_Grid_Engine) Qmaster Service"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6445",
		Description: `
Sun Grid Engine Execution Service

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6463–6472",
		Description: `
"Discord (https://en.wikipedia.org/wiki/Discord_(software)) RPC"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6464",
		Description: `
"Port assignment for medical device communication in accordance to IEEE 11073-20701 (https://en.wikipedia.org/wiki/ISO/IEEE_11073)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6502",
		Description: `
Netop Remote Control

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6513",
		Description: `
"NETCONF (https://en.wikipedia.org/wiki/NETCONF) over TLS (https://en.wikipedia.org/wiki/Transport_Layer_Security)"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6514",
		Description: `
"Syslog over TLS"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6515",
		Description: `
"Elipse (https://en.wikipedia.org/wiki/Elipse_Software) RPC Protocol (REC)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6516",
		Description: `
"Windows Admin Center (https://en.wikipedia.org/wiki/Windows_Admin_Center)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "6543",
		Description: `
"Pylons project#Pyramid (https://en.wikipedia.org/wiki/Pylons_project#Pyramid) Default Pylons Pyramid web service port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6556",
		Description: `
"Check MK (https://en.wikipedia.org/wiki/Check_MK) Agent"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6566",
		Description: `
"SANE (https://en.wikipedia.org/wiki/Scanner_Access_Now_Easy) (Scanner Access Now Easy)—SANE network scanner daemon"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6560–6561",
		Description: `
"Speech-Dispatcher (/w/index.php?title=Speech-Dispatcher&action=edit&redlink=1) daemon"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6571",
		Description: `
"Windows Live FolderShare (https://en.wikipedia.org/wiki/Windows_Live_FolderShare) client"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "6600",
		Description: `
"Microsoft Hyper-V (https://en.wikipedia.org/wiki/Hyper-V) Live"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified


"Music Player Daemon (https://en.wikipedia.org/wiki/Music_Player_Daemon) (MPD)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6601",
		Description: `
"Microsoft Forefront Threat Management Gateway (https://en.wikipedia.org/wiki/Microsoft_Forefront_Threat_Management_Gateway)"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6602",
		Description: `
Microsoft Windows WSS Communication

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6619",
		Description: `
"odette-ftps, Odette File Transfer Protocol (https://en.wikipedia.org/wiki/OFTP) (OFTP (https://en.wikipedia.org/wiki/OFTP)) over TLS (https://en.wikipedia.org/wiki/Transport_Layer_Security)/SSL (https://en.wikipedia.org/wiki/Secure_Sockets_Layer)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6622",
		Description: `
Multicast FTP

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6653",
		Description: `
"OpenFlow (https://en.wikipedia.org/wiki/OpenFlow)"

IANA Status - Official
TCP         - Yes
UDP         - Assigned

`,
	},
	{
		Name: "6660–6664",
		Description: `
"Internet Relay Chat (https://en.wikipedia.org/wiki/Internet_Relay_Chat) (IRC)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6665–6669",
		Description: `
"Internet Relay Chat (https://en.wikipedia.org/wiki/Internet_Relay_Chat) (IRC)"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6679",
		Description: `
Osorno Automation Protocol (OSAUT)

IANA Status - Official
TCP         - Yes
UDP         - Yes


"IRC (https://en.wikipedia.org/wiki/Internet_Relay_Chat) SSL (https://en.wikipedia.org/wiki/Secure_Sockets_Layer) (Secure Internet Relay Chat)—often used"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6690",
		Description: `
Synology Cloud station

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6697",
		Description: `
"IRC (https://en.wikipedia.org/wiki/Internet_Relay_Chat) SSL (https://en.wikipedia.org/wiki/Secure_Sockets_Layer) (Secure Internet Relay Chat)—often used"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6699",
		Description: `
"WinMX (https://en.wikipedia.org/wiki/WinMX) (see also 6257)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6715",
		Description: `
AberMUD and derivatives default port

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6771",
		Description: `
"BitTorrent Local Peer Discovery (https://en.wikipedia.org/wiki/Local_Peer_Discovery)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "6783–6785",
		Description: `
"Splashtop Remote (https://en.wikipedia.org/wiki/Splashtop_Remote) server broadcast"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6881–6887",
		Description: `
"BitTorrent (https://en.wikipedia.org/wiki/BitTorrent_(protocol)) part of full range of ports used most often"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6888",
		Description: `
MUSE

IANA Status - Official
TCP         - Yes
UDP         - Yes


BitTorrent part of full range of ports used most often

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6889–6890",
		Description: `
BitTorrent part of full range of ports used most often

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6891–6900",
		Description: `
BitTorrent part of full range of ports used most often

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6891–6900",
		Description: `
"Windows Live Messenger (https://en.wikipedia.org/wiki/Windows_Live_Messenger) (File transfer)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6901",
		Description: `
"Windows Live Messenger (https://en.wikipedia.org/wiki/Windows_Live_Messenger) (Voice)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes


BitTorrent part of full range of ports used most often

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6902–6968",
		Description: `
BitTorrent part of full range of ports used most often

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "6969",
		Description: `
acmsoda

IANA Status - Official
TCP         - Yes
UDP         - Yes


BitTorrent tracker

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "6970–6999",
		Description: `
BitTorrent part of full range of ports used most often

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes


"QuickTime Streaming Server (https://en.wikipedia.org/wiki/QuickTime_Streaming_Server)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "7000",
		Description: `
"Default for Vuze (https://en.wikipedia.org/wiki/Vuze)'s built-in HTTPS (https://en.wikipedia.org/wiki/HTTPS) Bittorrent Tracker (https://en.wikipedia.org/wiki/Bittorrent_Tracker)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Avira (https://en.wikipedia.org/wiki/Avira) Server Management Console"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7001",
		Description: `
Avira Server Management Console

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Default for BEA (https://en.wikipedia.org/wiki/BEA_Systems) WebLogic Server (https://en.wikipedia.org/wiki/WebLogic)'s HTTP (https://en.wikipedia.org/wiki/HTTP) server, though often changed during installation"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7002",
		Description: `
Default for BEA WebLogic Server's HTTPS server, though often changed during installation

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7005",
		Description: `
"Default for BMC Software (https://en.wikipedia.org/wiki/BMC_Software) Control-M/Server (https://en.wikipedia.org/wiki/BMC_Control-M) and Control-M/Agent for Agent-to-Server, though often changed during installation"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7006",
		Description: `
Default for BMC Software Control-M/Server and Control-M/Agent for Server-to-Agent, though often changed during installation

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7010",
		Description: `
"Default for Cisco AON AMC (AON Management Console)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7022",
		Description: `
"Database mirroring endpoints"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7023",
		Description: `
Bryan Wilcutt T2-NMCS Protocol for SatCom Modems

IANA Status - Official
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "7025",
		Description: `
"Zimbra LMTP (https://en.wikipedia.org/wiki/Local_Mail_Transfer_Protocol) [mailbox]—local mail delivery"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7047",
		Description: `
"Zimbra (https://en.wikipedia.org/wiki/Zimbra) conversion server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7070",
		Description: `
"Real Time Streaming Protocol (https://en.wikipedia.org/wiki/Real_Time_Streaming_Protocol) (RTSP), used by QuickTime Streaming Server (https://en.wikipedia.org/wiki/QuickTime_Streaming_Server). TCP is used by default, UDP is used as an alternate."

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes/No

`,
	},
	{
		Name: "7133",
		Description: `
"Enemy Territory: Quake Wars (https://en.wikipedia.org/wiki/Enemy_Territory:_Quake_Wars)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7144",
		Description: `
"Peercast"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7145",
		Description: `
"Peercast"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7171",
		Description: `
"Tibia (https://en.wikipedia.org/wiki/Tibia_(computer_game))"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7262",
		Description: `
CNAP (Calypso Network Access Protocol)

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "7272",
		Description: `
WatchMe - WatchMe Monitoring

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "7306",
		Description: `
"Zimbra mysql [mailbox]"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7307",
		Description: `
"Zimbra mysql [logger]"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7312",
		Description: `
"Sibelius (https://en.wikipedia.org/wiki/Sibelius_notation_program) License Server"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "7396",
		Description: `
"Web control interface for Folding@home v7.3.6 (https://en.wikipedia.org/wiki/Folding@home#V7) and later"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7400",
		Description: `
"RTPS (Real Time Publish Subscribe) DDS (https://en.wikipedia.org/wiki/Data_Distribution_Service) Discovery"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "7401",
		Description: `
"RTPS (Real Time Publish Subscribe) DDS (https://en.wikipedia.org/wiki/Data_Distribution_Service) User-Traffic"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "7402",
		Description: `
"RTPS (Real Time Publish Subscribe) DDS (https://en.wikipedia.org/wiki/Data_Distribution_Service) Meta-Traffic"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "7471",
		Description: `
Stateless Transport Tunneling (STT)

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7473",
		Description: `
"Rise: The Vieneo Province (https://en.wikipedia.org/wiki/Rise:_The_Vieneo_Province)"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7474",
		Description: `
"Neo4J Server webadmin"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7478",
		Description: `
"Default port used by Open iT (https://en.wikipedia.org/wiki/Open_iT) Server."

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7542",
		Description: `
"Saratoga file transfer protocol"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "7547",
		Description: `
"CPE WAN Management Protocol (CWMP) Technical Report 069 (https://en.wikipedia.org/wiki/TR-069)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "7575",
		Description: `
"Populous: The Beginning (https://en.wikipedia.org/wiki/Populous:_The_Beginning) server"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "7624",
		Description: `
"Instrument Neutral Distributed Interface (https://en.wikipedia.org/wiki/Instrument_Neutral_Distributed_Interface)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "7631",
		Description: `
ERLPhase

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7634",
		Description: `
hddtemp—Utility to monitor hard drive temperature

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7652–7654",
		Description: `
"I2P (https://en.wikipedia.org/wiki/I2P) anonymizing overlay network"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7655",
		Description: `
I2P SAM Bridge Socket API

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "7656–7660",
		Description: `
I2P anonymizing overlay network

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7670",
		Description: `
"BrettspielWelt (https://en.wikipedia.org/wiki/BrettspielWelt) BSW Boardgame Portal"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7680",
		Description: `
"Delivery Optimization for Windows 10 (https://en.wikipedia.org/wiki/Windows_10)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7687",
		Description: `
"Bolt (https://en.wikipedia.org/wiki/Bolt_(network_protocol)) database connection"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7707–7708",
		Description: `
"Killing Floor (https://en.wikipedia.org/wiki/Killing_Floor_(2009_video_game))"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "7717",
		Description: `
Killing Floor

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "7777",
		Description: `
"iChat (https://en.wikipedia.org/wiki/IChat) server file transfer proxy"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Oracle Cluster File System 2"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Windows backdoor program tini.exe default"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Just Cause 2: Multiplayer Mod Server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Terraria (https://en.wikipedia.org/wiki/Terraria) default server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


San Andreas Multiplayer (SA-MP) default port server

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified


SCP: Secret Laboratory Multiplayer Server

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "7777–7788",
		Description: `
"Unreal Tournament series default server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "7831",
		Description: `
"Default used by Smartlaunch Internet Cafe Administration software"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7880",
		Description: `
"PowerSchool Gradebook Server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "7890",
		Description: `
Default that will be used by the iControl Internet Cafe Suite Administration software

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7915",
		Description: `
"Default for YSFlight server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7935",
		Description: `
"Fixed port used for Adobe Flash Debug Player to communicate with a debugger (Flash IDE, Flex Builder or fdb)."

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "7946",
		Description: `
"Docker Swarm communication among nodes"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "7990",
		Description: `
"Atlassian Bitbucket (https://en.wikipedia.org/wiki/Bitbucket) (default port)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8000",
		Description: `
"Commonly used for Internet radio streams such as SHOUTcast (https://en.wikipedia.org/wiki/SHOUTcast), Icecast (https://en.wikipedia.org/wiki/Icecast) and iTunes Radio (https://en.wikipedia.org/wiki/ITunes_Radio)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"DynamoDB (https://en.wikipedia.org/wiki/DynamoDB) Local"

IANA Status - Unofficial
TCP         - ?
UDP         - ?


"Django (https://en.wikipedia.org/wiki/Django_(web_framework)) Development Webserver"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8005",
		Description: `
"Tomcat (https://en.wikipedia.org/wiki/Apache_Tomcat) remote shutdown"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8006",
		Description: `
"Quest AppAssure (https://en.wikipedia.org/wiki/Quest_AppAssure) 5 API"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Proxmox Virtual Environment (https://en.wikipedia.org/wiki/Proxmox_Virtual_Environment) admin web interface"

IANA Status - Unofficial
TCP         - Yes
UDP         - No

`,
	},
	{
		Name: "8007",
		Description: `
"Quest AppAssure 5 Engine"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8008",
		Description: `
"Alternative port for HTTP (https://en.wikipedia.org/wiki/HTTP). See also ports 80 and 8080."

IANA Status - Official
TCP         - Yes
UDP         - Yes


"IBM HTTP Server (https://en.wikipedia.org/wiki/IBM_HTTP_Server) administration default"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"iCal (https://en.wikipedia.org/wiki/ICal), a calendar application by Apple (https://en.wikipedia.org/wiki/Apple,_Inc.)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Matrix (https://en.wikipedia.org/wiki/Matrix_(protocol)) homeserver federation over HTTP"

IANA Status - Unofficial
TCP         - Yes
UDP         - ?

`,
	},
	{
		Name: "8009",
		Description: `
"Apache JServ Protocol (https://en.wikipedia.org/wiki/Apache_JServ_Protocol) (ajp13)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8010",
		Description: `
"Buildbot (https://en.wikipedia.org/wiki/Buildbot) Web status page"

IANA Status - Unofficial
TCP         - Yes
UDP         - ?

`,
	},
	{
		Name: "8042",
		Description: `
"Orthanc (https://en.wikipedia.org/wiki/Orthanc_(software)) – REST API over HTTP"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8069",
		Description: `
"OpenERP (https://en.wikipedia.org/wiki/OpenERP) 5.0 XML-RPC protocol"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8070",
		Description: `
"OpenERP 5.0 NET-RPC protocol"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8074",
		Description: `
"Gadu-Gadu (https://en.wikipedia.org/wiki/Gadu-Gadu)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "8075",
		Description: `
"Killing Floor web administration interface"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8080",
		Description: `
"Alternative port for HTTP (https://en.wikipedia.org/wiki/HTTP). See also ports 80 and 8008."

IANA Status - Official
TCP         - Yes
UDP         - Yes


"Apache Tomcat (https://en.wikipedia.org/wiki/Apache_Tomcat)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Atlassian JIRA (https://en.wikipedia.org/wiki/Jira_(software)) applications"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8088",
		Description: `
"Asterisk (https://en.wikipedia.org/wiki/Asterisk_(PBX)) management access via HTTP"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8089",
		Description: `
"Splunk (https://en.wikipedia.org/wiki/Splunk) daemon management"

IANA Status - Unofficial
TCP         - Yes
UDP         - No


"Fritz!Box (https://en.wikipedia.org/wiki/Fritz!Box) automatic TR-069 (https://en.wikipedia.org/wiki/TR-069) configuration"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8090",
		Description: `
"Atlassian Confluence (https://en.wikipedia.org/wiki/Atlassian_Confluence)"

IANA Status - Unofficial
TCP         - ?
UDP         - ?


"Coral Content Distribution Network (https://en.wikipedia.org/wiki/Coral_Content_Distribution_Network) (legacy; 80 and 8080 now supported)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Matrix identity server"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8091",
		Description: `
"CouchBase (https://en.wikipedia.org/wiki/CouchBase) web administration"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8092",
		Description: `
"CouchBase API"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8111",
		Description: `
"JOSM (https://en.wikipedia.org/wiki/JOSM) Remote Control"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8112",
		Description: `
PAC Pacifica Coin

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8116",
		Description: `
"Check Point (https://en.wikipedia.org/wiki/Check_Point) Cluster Control Protocol"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "8118",
		Description: `
"Privoxy (https://en.wikipedia.org/wiki/Privoxy)—advertisement-filtering Web proxy"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8123",
		Description: `
"Polipo (https://en.wikipedia.org/wiki/Polipo) Web proxy"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8139",
		Description: `
"Puppet (software) (https://en.wikipedia.org/wiki/Puppet_(software)) Client agent"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8140",
		Description: `
Puppet (software) Master server

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8172",
		Description: `
"Microsoft (https://en.wikipedia.org/wiki/Microsoft) Remote Administration for IIS Manager"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8184",
		Description: `
"NCSA Brown Dog (https://en.wikipedia.org/wiki/NCSA_Brown_Dog) Data Access Proxy"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8194–8195",
		Description: `
"Bloomberg Terminal (https://en.wikipedia.org/wiki/Bloomberg_Terminal)"

IANA Status - Official
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8200",
		Description: `
"GoToMyPC (https://en.wikipedia.org/wiki/GoToMyPC)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


MiniDLNA media server Web Interface

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8222",
		Description: `
"VMware (https://en.wikipedia.org/wiki/VMware) VI Web Access via HTTP"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8243",
		Description: `
"HTTPS (https://en.wikipedia.org/wiki/HTTPS) listener for Apache Synapse (https://en.wikipedia.org/wiki/Apache_Synapse)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "8245",
		Description: `
"Dynamic DNS (https://en.wikipedia.org/wiki/Dynamic_DNS) for at least No-IP (https://en.wikipedia.org/wiki/No-IP) and DyDNS"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8280",
		Description: `
"HTTP (https://en.wikipedia.org/wiki/HTTP) listener for Apache Synapse (https://en.wikipedia.org/wiki/Apache_Synapse)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "8281",
		Description: `
HTTP Listener for Gatecraft Plugin

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8291",
		Description: `
"Winbox—Default on a MikroTik RouterOS for a Windows application used to administer MikroTik RouterOS"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8303",
		Description: `
"Teeworlds (https://en.wikipedia.org/wiki/Teeworlds) Server"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "8332",
		Description: `
"Bitcoin (https://en.wikipedia.org/wiki/Bitcoin) JSON-RPC (https://en.wikipedia.org/wiki/JSON-RPC) server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8333",
		Description: `
"Bitcoin (https://en.wikipedia.org/wiki/Bitcoin)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"VMware (https://en.wikipedia.org/wiki/VMware) VI Web Access via HTTPS"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8337",
		Description: `
"VisualSVN Distributed File System Service (VDFS)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8384",
		Description: `
"Syncthing (https://en.wikipedia.org/wiki/Syncthing) web GUI"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8388",
		Description: `
"Shadowsocks (https://en.wikipedia.org/wiki/Shadowsocks) proxy server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8443",
		Description: `
"SW Soft Plesk (https://en.wikipedia.org/wiki/SW_Soft_Plesk) Control Panel"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Apache Tomcat (https://en.wikipedia.org/wiki/Apache_Tomcat) SSL"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Promise (https://en.wikipedia.org/wiki/Promise_Technology) WebPAM SSL"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"iCal (https://en.wikipedia.org/wiki/ICal) over SSL (https://en.wikipedia.org/wiki/Secure_Socket_Layer)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"MineOs (/w/index.php?title=MineOs&action=edit&redlink=1) WebUi"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8444",
		Description: `
"Bitmessage (https://en.wikipedia.org/wiki/Bitmessage)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8448",
		Description: `
"Matrix homeserver federation over HTTPS"

IANA Status - Unofficial
TCP         - Yes
UDP         - ?

`,
	},
	{
		Name: "8484",
		Description: `
"MapleStory (https://en.wikipedia.org/wiki/MapleStory) Login Server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8500",
		Description: `
"Adobe ColdFusion (https://en.wikipedia.org/wiki/Adobe_ColdFusion) built-in web server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8530",
		Description: `
"Windows Server Update Services (https://en.wikipedia.org/wiki/Windows_Server_Update_Services) over HTTP (https://en.wikipedia.org/wiki/HTTP)"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8531",
		Description: `
"Windows Server Update Services over HTTPS (https://en.wikipedia.org/wiki/HTTPS)"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8580",
		Description: `
"Freegate (https://en.wikipedia.org/wiki/Freegate), an Internet anonymizer (https://en.wikipedia.org/wiki/Anonymizer) and proxy tool"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8629",
		Description: `
"Tibero database (https://en.wikipedia.org/wiki/Tibero)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8642",
		Description: `
"Lotus Notes Traveler (https://en.wikipedia.org/wiki/IBM_Lotus_Notes_Traveler) auto synchronization for Windows Mobile and Nokia devices"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8691",
		Description: `
"Ultra Fractal (https://en.wikipedia.org/wiki/Ultra_Fractal), a fractal (https://en.wikipedia.org/wiki/Fractal) generation and rendering software application (https://en.wikipedia.org/wiki/Graphic_art_software) – distributed calculations over networked computers"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "8765",
		Description: `
"Default port of a local GUN relay peer that the Internet Archive (https://en.wikipedia.org/wiki/Internet_Archive) and others use as a decentralized mirror for censorship resistance."

IANA Status - Unofficial
TCP         - Yes
UDP         - ?

`,
	},
	{
		Name: "8767",
		Description: `
"Voice channel of TeamSpeak 2 (https://en.wikipedia.org/wiki/TeamSpeak_2), a proprietary Voice over IP (https://en.wikipedia.org/wiki/Voice_over_IP) protocol targeted at gamers (https://en.wikipedia.org/wiki/Gamer)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "8834",
		Description: `
"Nessus (https://en.wikipedia.org/wiki/Nessus_(software)), a vulnerability scanner (https://en.wikipedia.org/wiki/Vulnerability_scanner) – remote XML-RPC (https://en.wikipedia.org/wiki/XML-RPC) web server"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8840",
		Description: `
"Opera Unite (/w/index.php?title=Opera_Unite&action=edit&redlink=1), an extensible framework (https://en.wikipedia.org/wiki/Software_framework) for web applications"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8880",
		Description: `
"Alternate port of CDDB (https://en.wikipedia.org/wiki/CDDB) (Compact Disc Database) protocol, used to look up audio CD (compact disc (https://en.wikipedia.org/wiki/Compact_disc)) information over the Internet (https://en.wikipedia.org/wiki/Internet). See also port 888."

IANA Status - Official
TCP         - Yes
UDP         - Unspecified


"IBM WebSphere Application Server (https://en.wikipedia.org/wiki/IBM_WebSphere_Application_Server) SOAP (https://en.wikipedia.org/wiki/SOAP) connector"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8883",
		Description: `
"Secure MQTT (https://en.wikipedia.org/wiki/MQTT) (MQTT over TLS)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "8887",
		Description: `
"HyperVM (https://en.wikipedia.org/wiki/HyperVM) over HTTP"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8888",
		Description: `
"HyperVM over HTTPS (https://en.wikipedia.org/wiki/HTTPS)"

IANA Status - Unofficial
TCP         - ?
UDP         - ?


"Freenet (https://en.wikipedia.org/wiki/Freenet) web UI (localhost only)"

IANA Status - Unofficial
TCP         - ?
UDP         - Yes


"Default for IPython (https://en.wikipedia.org/wiki/IPython) / Jupyter (https://en.wikipedia.org/wiki/Jupyter) notebook dashboards"

IANA Status - Unofficial
TCP         - ?
UDP         - ?


"MAMP (https://en.wikipedia.org/wiki/MAMP)"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8889",
		Description: `
"MAMP"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8983",
		Description: `
"Apache Solr (https://en.wikipedia.org/wiki/Apache_Solr)"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8997",
		Description: `
"Alternate port for I2P (https://en.wikipedia.org/wiki/I2P) Monotone Proxy"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8998",
		Description: `
"I2P Monotone Proxy"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "8999",
		Description: `
"Alternate port for I2P Monotone Proxy"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "9000",
		Description: `
"SonarQube (https://en.wikipedia.org/wiki/SonarQube) Web Server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"DBGp (https://en.wikipedia.org/wiki/DBGp)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"SqueezeCenter (https://en.wikipedia.org/wiki/SqueezeCenter) web server & streaming"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"UDPCast (https://en.wikipedia.org/wiki/UDPCast)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified


"Play! Framework web server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Hadoop (https://en.wikipedia.org/wiki/Hadoop_Distributed_File_System) NameNode default port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"PHP-FPM (https://en.wikipedia.org/wiki/PHP-FPM) default port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"QBittorrent (https://en.wikipedia.org/wiki/QBittorrent)'s embedded torrent tracker default port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9001",
		Description: `
"ETL Service Manager"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"Microsoft SharePoint (https://en.wikipedia.org/wiki/Microsoft_SharePoint) authoring environment"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified


"cisco-xremote router configuration"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified


"Tor (https://en.wikipedia.org/wiki/Tor_(anonymity_network)) network default"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified


"DBGp (https://en.wikipedia.org/wiki/DBGp) Proxy"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"HSQLDB (https://en.wikipedia.org/wiki/HSQLDB) default port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9002",
		Description: `
Newforma Server comms

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "9006",
		Description: `
De-Commissioned Port

IANA Status - Official
TCP         - Unspecified
UDP         - Unspecified


"Tomcat (https://en.wikipedia.org/wiki/Apache_Tomcat) in standalone mode"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9030",
		Description: `
"Tor (https://en.wikipedia.org/wiki/Tor_(anonymity_network)) often used"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9042",
		Description: `
"Apache Cassandra (https://en.wikipedia.org/wiki/Apache_Cassandra) native protocol clients"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9043",
		Description: `
"WebSphere Application Server (https://en.wikipedia.org/wiki/WebSphere_Application_Server) Administration Console secure"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9050–9051",
		Description: `
"Tor (https://en.wikipedia.org/wiki/Tor_(anonymity_network)) (SOCKS-5 proxy client)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9060",
		Description: `
"WebSphere Application Server (https://en.wikipedia.org/wiki/WebSphere_Application_Server) Administration Console"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9080",
		Description: `
"glrpc, Groove (https://en.wikipedia.org/wiki/Microsoft_Groove) Collaboration software (https://en.wikipedia.org/wiki/Collaboration_software) GLRPC"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"WebSphere Application Server (https://en.wikipedia.org/wiki/WebSphere_Application_Server) HTTP (https://en.wikipedia.org/wiki/HTTP) Transport (port 1) default (https://en.wikipedia.org/wiki/Default_(computer_science))"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


Remote Potato by FatAttitude, Windows Media Center addon

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


ServerWMC, Windows Media Center addon

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9090",
		Description: `
"Prometheus (https://en.wikipedia.org/wiki/Prometheus_(software)) metrics server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Openfire (https://en.wikipedia.org/wiki/Openfire) Administration Console"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"SqueezeCenter (https://en.wikipedia.org/wiki/SqueezeCenter) control (CLI)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Cherokee (https://en.wikipedia.org/wiki/Cherokee_(web_server)) Admin Panel"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9091",
		Description: `
"Openfire (https://en.wikipedia.org/wiki/Openfire) Administration Console (SSL Secured)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Transmission (BitTorrent client) (https://en.wikipedia.org/wiki/Transmission_(BitTorrent_client)) Web Interface"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9092",
		Description: `
"H2 (DBMS) (https://en.wikipedia.org/wiki/H2_(DBMS)) Database Server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Apache Kafka (https://en.wikipedia.org/wiki/Apache_Kafka) A Distributed Streaming Platform"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9100",
		Description: `
"PDL (https://en.wikipedia.org/wiki/Page_description_language) Data Stream, used for printing to certain network printers"

IANA Status - Official
TCP         - Yes
UDP         - Assigned

`,
	},
	{
		Name: "9101",
		Description: `
"Bacula (https://en.wikipedia.org/wiki/Bacula) Director"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "9102",
		Description: `
"Bacula (https://en.wikipedia.org/wiki/Bacula) File Daemon"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "9103",
		Description: `
"Bacula (https://en.wikipedia.org/wiki/Bacula) Storage Daemon"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "9119",
		Description: `
"MXit (https://en.wikipedia.org/wiki/MXit) Instant Messenger"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "9150",
		Description: `
"Tor (https://en.wikipedia.org/wiki/Tor_(anonymity_network)) Browser"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9191",
		Description: `
Sierra Wireless Airlink

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9199",
		Description: `
Avtex LLC—qStats

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9200",
		Description: `
"Elasticsearch—default Elasticsearch port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9217",
		Description: `
iPass Platform Service

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9293",
		Description: `
"Sony PlayStation RemotePlay"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9300",
		Description: `
"IBM Cognos BI (https://en.wikipedia.org/wiki/IBM_Cognos_Business_Intelligence)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9303",
		Description: `
"D-Link Shareport (/w/index.php?title=D-Link_Shareport&action=edit&redlink=1) Share storage and MFP printers"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "9306",
		Description: `
"Sphinx (https://en.wikipedia.org/wiki/Sphinx_(search_engine)) Native API"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9309",
		Description: `
"Sony PlayStation Vita Host Collaboration WiFi Data Transfer"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "9312",
		Description: `
"Sphinx (https://en.wikipedia.org/wiki/Sphinx_(search_engine)) SphinxQL"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9332",
		Description: `
"Litecoin (https://en.wikipedia.org/wiki/Litecoin) JSON-RPC (https://en.wikipedia.org/wiki/JSON-RPC) server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9333",
		Description: `
"Litecoin (https://en.wikipedia.org/wiki/Litecoin)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9339",
		Description: `
Used by all Supercell games such as Brawl Stars and Clash of Clans, mobile freemium strategy video games

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9389",
		Description: `
"adws, Microsoft (https://en.wikipedia.org/wiki/Microsoft) AD DS (https://en.wikipedia.org/wiki/AD_DS) Web Services, Powershell (https://en.wikipedia.org/wiki/Powershell) uses this port"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "9392",
		Description: `
OpenVAS Greenbone Security Assistant web interface

IANA Status - Unofficial
TCP         - Yes
UDP         - ?

`,
	},
	{
		Name: "9418",
		Description: `
"git, Git (https://en.wikipedia.org/wiki/Git_(software)) pack transfer service"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "9419",
		Description: `
"MooseFS (https://en.wikipedia.org/wiki/Moose_File_System) distributed file system – master control port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9420",
		Description: `
"MooseFS distributed file system – master command port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9421",
		Description: `
"MooseFS distributed file system – master client port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9422",
		Description: `
"MooseFS distributed file system – Chunkservers"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9425",
		Description: `
"MooseFS distributed file system – CGI server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9443",
		Description: `
"VMware (https://en.wikipedia.org/wiki/VMware) Websense Triton console (HTTPS port used for accessing and administrating a vCenter Server via the Web Management Interface)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"NCSA Brown Dog (https://en.wikipedia.org/wiki/NCSA_Brown_Dog) Data Tilling Service"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9535",
		Description: `
"mngsuite, LANDesk (https://en.wikipedia.org/wiki/Landesk) Management Suite Remote Control"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "9536",
		Description: `
"laes-bf, IP Fabrics (https://en.wikipedia.org/wiki/IP_Fabrics) Surveillance buffering function"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "9600",
		Description: `
"Factory Interface Network Service (https://en.wikipedia.org/wiki/Factory_Interface_Network_Service) (FINS), a network protocol used by Omron (https://en.wikipedia.org/wiki/Omron) programmable logic controllers (https://en.wikipedia.org/wiki/Programmable_logic_controller)"

IANA Status - Unofficial
TCP         - No
UDP         - Yes

`,
	},
	{
		Name: "9675",
		Description: `
"Spiceworks (https://en.wikipedia.org/wiki/Spiceworks) Desktop, IT Helpdesk Software"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "9676",
		Description: `
"Spiceworks (https://en.wikipedia.org/wiki/Spiceworks) Desktop, IT Helpdesk Software"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "9695",
		Description: `
"Content centric networking (https://en.wikipedia.org/wiki/Content_centric_networking) (CCN, CCNx)"

IANA Status - Official
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "9785",
		Description: `
"Viber (https://en.wikipedia.org/wiki/Viber)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "9800",
		Description: `
"WebDAV (https://en.wikipedia.org/wiki/WebDAV) Source"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"WebCT (https://en.wikipedia.org/wiki/WebCT) e-learning portal"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "9875",
		Description: `
"Club Penguin (https://en.wikipedia.org/wiki/Club_Penguin) Disney online game for kids"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9898",
		Description: `
"Tripwire (https://en.wikipedia.org/wiki/Tripwire_(software))—File Integrity Monitoring Software"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9899",
		Description: `
"SCTP (https://en.wikipedia.org/wiki/Stream_Control_Transmission_Protocol) tunneling (port number used in SCTP packets encapsulated in UDP, RFC 6951 (https://tools.ietf.org/html/rfc6951))"

IANA Status - Official
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "9901",
		Description: `
Banana for Apache Solr

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9981",
		Description: `
"Tvheadend (https://en.wikipedia.org/wiki/Tvheadend) HTTP server (web interface)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9982",
		Description: `
"Tvheadend (https://en.wikipedia.org/wiki/Tvheadend) HTSP server (Streaming protocol)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9987",
		Description: `
"TeamSpeak (https://en.wikipedia.org/wiki/TeamSpeak) 3 server default (voice) port (for the conflicting service see the IANA list)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "9993",
		Description: `
"ZeroTier (https://en.wikipedia.org/wiki/ZeroTier) Default port for ZeroTier"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "9997",
		Description: `
"Splunk (https://en.wikipedia.org/wiki/Splunk) port for communication between the forwarders and indexers"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "9999",
		Description: `
"Urchin (https://en.wikipedia.org/wiki/Urchin_Software_Corporation) Web Analytics"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "10000",
		Description: `
Network Data Management Protocol

IANA Status - Official
TCP         - Yes
UDP         - Yes


"BackupExec (https://en.wikipedia.org/wiki/BackupExec)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified


"Webmin (https://en.wikipedia.org/wiki/Webmin), Web-based Unix/Linux system administration tool (default port)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "10000–20000",
		Description: `
"Used on VoIP (https://en.wikipedia.org/wiki/Voice_over_IP) networks for receiving and transmitting voice telephony traffic which includes Google Voice (https://en.wikipedia.org/wiki/Google_Voice) via the OBiTalk (https://en.wikipedia.org/wiki/Obihai_Technology) ATA (https://en.wikipedia.org/wiki/Analog_telephone_adapter) devices as well as on the MagicJack (https://en.wikipedia.org/wiki/MagicJack) and Vonage (https://en.wikipedia.org/wiki/Vonage) ATA network devices."

IANA Status - Unofficial
TCP         - No
UDP         - Yes

`,
	},
	{
		Name: "10001",
		Description: `
Ubiquiti UniFi access points broadcast to 255.255.255.255:10001 (UDP) to locate the controller(s)

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "10009",
		Description: `
"CrossFire (https://en.wikipedia.org/wiki/CrossFire_(video_game)), a multiplayer online First Person Shooter"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "10011",
		Description: `
"Teamspeak3 Chat Server"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "10024",
		Description: `
"Zimbra smtp [mta]—to amavis from postfix"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "10025",
		Description: `
"Zimbra smtp [mta]—back to postfix from amavis"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "10042",
		Description: `
"Mathoid (https://www.mediawiki.org/wiki/Mathoid) server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "10050",
		Description: `
"Zabbix (https://en.wikipedia.org/wiki/Zabbix) agent"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "10051",
		Description: `
"Zabbix (https://en.wikipedia.org/wiki/Zabbix) trapper"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "10110",
		Description: `
NMEA 0183 Navigational Data. Transport of NMEA 0183 sentences over TCP or UDP

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "10172",
		Description: `
"Intuit Quickbooks (https://en.wikipedia.org/wiki/Quickbooks) client"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "10200",
		Description: `
"FRISK Software International (https://en.wikipedia.org/wiki/FRISK_Software_International)'s fpscand virus scanning daemon for Unix platforms"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"FRISK Software International's f-protd virus scanning daemon for Unix platforms"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "10201–10204",
		Description: `
"FRISK Software International's f-protd virus scanning daemon for Unix platforms"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "10212",
		Description: `
"GE Intelligent Platforms Proficy HMI/SCADA – CIMPLICITY WebView"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "10308",
		Description: `
"Lock On: Modern Air Combat (https://en.wikipedia.org/wiki/Lock_On:_Modern_Air_Combat)"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "10480",
		Description: `
"SWAT 4 (https://en.wikipedia.org/wiki/SWAT_4) Dedicated Server"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "10505",
		Description: `
"BlueStacks (android simulator) broadcast"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "10514",
		Description: `
TLS-enabled Rsyslog (default by convention)

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "10800",
		Description: `
"Touhou (https://en.wikipedia.org/wiki/Touhou) fight games (Immaterial and Missing Power (https://en.wikipedia.org/wiki/Immaterial_and_Missing_Power), Scarlet Weather Rhapsody (https://en.wikipedia.org/wiki/Scarlet_Weather_Rhapsody), Hisoutensoku (https://en.wikipedia.org/wiki/Touhou_His%C5%8Dtensoku), Hopeless Masquerade (https://en.wikipedia.org/wiki/Hopeless_Masquerade) and Urban Legend in Limbo (https://en.wikipedia.org/wiki/Urban_Legend_in_Limbo))"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "10823",
		Description: `
"Farming Simulator 2011 (https://en.wikipedia.org/wiki/Farming_Simulator_2011)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "10891",
		Description: `
"Jungle Disk (this port is opened by the Jungle Disk Monitor service on the localhost)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "10933",
		Description: `
"Octopus Deploy Tentacle deployment agent"

IANA Status - Official
TCP         - Yes
UDP         - No

`,
	},
	{
		Name: "11001",
		Description: `
metasys ( Johnson Controls Metasys java AC control environment )

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "11111",
		Description: `
RiCcI, Remote Configuration Interface (Redhat Linux)

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "11112",
		Description: `
"ACR (https://en.wikipedia.org/wiki/American_College_of_Radiology)/NEMA (https://en.wikipedia.org/wiki/National_Electrical_Manufacturers_Association) Digital Imaging and Communications in Medicine (https://en.wikipedia.org/wiki/Digital_Imaging_and_Communications_in_Medicine) (DICOM)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "11211",
		Description: `
"memcached (https://en.wikipedia.org/wiki/Memcached)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "11214",
		Description: `
memcached incoming SSL proxy

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "11215",
		Description: `
memcached internal outgoing SSL proxy

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "11235",
		Description: `
"Savage: Battle for Newerth (https://en.wikipedia.org/wiki/Savage:_Battle_for_Newerth)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "11311",
		Description: `
"Robot Operating System (https://en.wikipedia.org/wiki/Robot_Operating_System) master"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "11371",
		Description: `
"OpenPGP (https://en.wikipedia.org/wiki/OpenPGP) HTTP key server (https://en.wikipedia.org/wiki/Key_server_(cryptographic))"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "11753",
		Description: `
"OpenRCT2 (https://en.wikipedia.org/wiki/OpenRCT2) multiplayer"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "12000",
		Description: `
"CubeForm (/w/index.php?title=CubeForm&action=edit&redlink=1), Multiplayer SandBox Game"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "12012",
		Description: `
"Audition Online Dance Battle (https://en.wikipedia.org/wiki/Audition_Online_Dance_Battle), Korea Server—Status/Version Check"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "12013",
		Description: `
Audition Online Dance Battle, Korea Server

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "12035",
		Description: `
"Second Life (https://en.wikipedia.org/wiki/Second_Life), used for server UDP in-bound"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "12043",
		Description: `
"Second Life, used for LSL HTTPS in-bound"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "12046",
		Description: `
"Second Life, used for LSL HTTP in-bound"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "12201",
		Description: `
"Graylog Extended Log Format (GELF)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "12222",
		Description: `
"Light Weight Access Point Protocol (LWAPP (https://en.wikipedia.org/wiki/LWAPP)) LWAPP data (RFC 5412 (https://tools.ietf.org/html/rfc5412))"

IANA Status - Official
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "12223",
		Description: `
"Light Weight Access Point Protocol (LWAPP (https://en.wikipedia.org/wiki/LWAPP)) LWAPP control (RFC 5412 (https://tools.ietf.org/html/rfc5412))"

IANA Status - Official
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "12345",
		Description: `
"Cube World (https://en.wikipedia.org/wiki/Cube_World)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes


"Little Fighter 2 (https://en.wikipedia.org/wiki/Little_Fighter_2)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"NetBus (https://en.wikipedia.org/wiki/NetBus) remote administration tool (often Trojan horse (https://en.wikipedia.org/wiki/Trojan_horse_(computing)))."

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "12443",
		Description: `
"IBM HMC (https://en.wikipedia.org/wiki/IBM_Hardware_Management_Console) web browser management access over HTTPS (https://en.wikipedia.org/wiki/HTTPS) instead of default port 443"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "12489",
		Description: `
NSClient/NSClient++/NC_Net (Nagios)

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "12975",
		Description: `
"LogMeIn (https://en.wikipedia.org/wiki/LogMeIn) Hamachi (https://en.wikipedia.org/wiki/Hamachi_(software)) (VPN tunnel software; also port 32976)—used to connect to Mediation Server (bibi.hamachi.cc); will attempt to use SSL (https://en.wikipedia.org/wiki/Secure_Sockets_Layer) (TCP port 443) if both 12975 & 32976 fail to connect"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "13000–13050",
		Description: `
"Second Life (https://en.wikipedia.org/wiki/Second_Life), used for server UDP in-bound"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "13008",
		Description: `
"CrossFire (https://en.wikipedia.org/wiki/CrossFire_(video_game)), a multiplayer online First Person Shooter"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "13075",
		Description: `
"Default for BMC Software (https://en.wikipedia.org/wiki/BMC_Software) Control-M/Enterprise Manager (https://en.wikipedia.org/wiki/BMC_Control-M) Corba communication, though often changed during installation"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "13400",
		Description: `
ISO 13400 Road vehicles — Diagnostic communication over Internet Protocol(DoIP)

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "13720",
		Description: `
"Symantec (https://en.wikipedia.org/wiki/NortonLifeLock) NetBackup (https://en.wikipedia.org/wiki/NetBackup)—bprd (formerly VERITAS (https://en.wikipedia.org/wiki/Veritas_Software))"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "13721",
		Description: `
Symantec NetBackup—bpdbm (formerly VERITAS)

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "13724",
		Description: `
Symantec Network Utility—vnetd (formerly VERITAS)

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "13782",
		Description: `
Symantec NetBackup—bpcd (formerly VERITAS)

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "13783",
		Description: `
Symantec VOPIED protocol (formerly VERITAS)

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "13785",
		Description: `
Symantec NetBackup Database—nbdb (formerly VERITAS)

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "13786",
		Description: `
Symantec nomdb (formerly VERITAS)

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "14550",
		Description: `
"MAVLink (https://en.wikipedia.org/wiki/MAVLink) Ground Station Port"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "14567",
		Description: `
"Battlefield 1942 (https://en.wikipedia.org/wiki/Battlefield_1942) and mods"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "14800",
		Description: `
"Age of Wonders III (https://en.wikipedia.org/wiki/Age_of_Wonders_III) p2p port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "15000",
		Description: `
"psyBNC (https://en.wikipedia.org/wiki/PsyBNC)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Wesnoth (https://en.wikipedia.org/wiki/Wesnoth)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Kaspersky Network Agent"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


Teltonika networks remote management system (RMS)

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "15009",
		Description: `
Teltonika networks remote management system (RMS)

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "15010",
		Description: `
Teltonika networks remote management system (RMS)

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "15441",
		Description: `
"ZeroNet (https://en.wikipedia.org/wiki/ZeroNet) fileserver"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "15567",
		Description: `
"Battlefield Vietnam (https://en.wikipedia.org/wiki/Battlefield_Vietnam) and mods"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "15345",
		Description: `
"XPilot (https://en.wikipedia.org/wiki/XPilot) Contact"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "15672",
		Description: `
"RabbitMQ (https://en.wikipedia.org/wiki/RabbitMQ) management plugin"

IANA Status - Unofficial
TCP         - Yes
UDP         - No

`,
	},
	{
		Name: "16000",
		Description: `
"Oracle WebCenter (https://en.wikipedia.org/wiki/Oracle_WebCenter) Content: Imaging (formerly known as Oracle Universal Content Management (https://en.wikipedia.org/wiki/Universal_Content_Management)). Port though often changed during installation"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"shroudBNC (https://en.wikipedia.org/wiki/ShroudBNC)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "16080",
		Description: `
"Mac OS X Server (https://en.wikipedia.org/wiki/Mac_OS_X_Server) Web (HTTP) service with performance cache"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "16200",
		Description: `
"Oracle WebCenter Content: Content Server (formerly known as Oracle Universal Content Management (https://en.wikipedia.org/wiki/Universal_Content_Management)). Port though often changed during installation"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "16225",
		Description: `
Oracle WebCenter Content: Content Server Web UI. Port though often changed during installation

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "16250",
		Description: `
"Oracle WebCenter Content: Inbound Refinery (formerly known as Oracle Universal Content Management (https://en.wikipedia.org/wiki/Universal_Content_Management)). Port though often changed during installation"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "16261",
		Description: `
"Project Zomboid (https://en.wikipedia.org/wiki/Project_Zomboid) multiplayer. Additional sequential ports used for each player connecting to server."

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "16300",
		Description: `
"Oracle WebCenter Content: Records Management (formerly known as Oracle Universal Records Management (/w/index.php?title=Universal_Records_Management&action=edit&redlink=1)). Port though often changed during installation"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "16384",
		Description: `
CISCO Default RTP MIN

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "16384-16403",
		Description: `
"Real-time Transport Protocol (https://en.wikipedia.org/wiki/Real-time_Transport_Protocol) (RTP), RTP Control Protocol (https://en.wikipedia.org/wiki/RTP_Control_Protocol) (RTCP), used by Apple (https://en.wikipedia.org/wiki/Apple_Inc.)'s iChat (https://en.wikipedia.org/wiki/IChat) for audio and video"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "16384-16387",
		Description: `
"Real-time Transport Protocol (RTP), RTP Control Protocol (RTCP), used by Apple's FaceTime (https://en.wikipedia.org/wiki/FaceTime) and Game Center (https://en.wikipedia.org/wiki/Game_Center)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "16393-16402",
		Description: `
"Real-time Transport Protocol (RTP), RTP Control Protocol (RTCP), used by Apple's FaceTime and Game Center"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "16403-16472",
		Description: `
"Real-time Transport Protocol (RTP), RTP Control Protocol (RTCP), used by Apple's Game Center"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "16400",
		Description: `
Oracle WebCenter Content: Capture (formerly known as Oracle Document Capture). Port though often changed during installation

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "16482",
		Description: `
CISCO Default RTP MAX

IANA Status - Official
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "16567",
		Description: `
"Battlefield 2 (https://en.wikipedia.org/wiki/Battlefield_2) and mods"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "17011",
		Description: `
"Worms (https://en.wikipedia.org/wiki/Worms_(series)) multiplayer"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "17224",
		Description: `
"Train Realtime Data Protocol (TRDP) Process Data, network protocol used in train communication."

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "17225",
		Description: `
"Train Realtime Data Protocol (TRDP) Message Data, network protocol used in train communication."

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "17333",
		Description: `
CS Server (CSMS), default binary protocol port

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "17475",
		Description: `
DMXControl 3 Network Broker

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "17500",
		Description: `
"Dropbox (https://en.wikipedia.org/wiki/Dropbox_(storage_provider)) LanSync Protocol (db-lsp); used to synchronize file catalogs between Dropbox clients on a local network."

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "18080",
		Description: `
"Monero (https://en.wikipedia.org/wiki/Monero_(cryptocurrency)) P2P network communications"

IANA Status - Unofficial
TCP         - Yes
UDP         - No

`,
	},
	{
		Name: "18081",
		Description: `
"Monero incoming RPC calls"

IANA Status - Unofficial
TCP         - Yes
UDP         - No

`,
	},
	{
		Name: "18091",
		Description: `
"memcached (https://en.wikipedia.org/wiki/Memcached) Internal REST HTTPS for SSL"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "18092",
		Description: `
memcached Internal CAPI HTTPS for SSL

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "18104",
		Description: `
RAD PDF Service

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "18200",
		Description: `
"Audition Online Dance Battle (https://en.wikipedia.org/wiki/Audition_Online_Dance_Battle), AsiaSoft Thailand Server status/version check"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "18201",
		Description: `
Audition Online Dance Battle, AsiaSoft Thailand Server

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "18206",
		Description: `
Audition Online Dance Battle, AsiaSoft Thailand Server FAM database

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "18300",
		Description: `
Audition Online Dance Battle, AsiaSoft SEA Server status/version check

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "18301",
		Description: `
Audition Online Dance Battle, AsiaSoft SEA Server

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "18306",
		Description: `
Audition Online Dance Battle, AsiaSoft SEA Server FAM database

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "18333",
		Description: `
"Bitcoin (https://en.wikipedia.org/wiki/Bitcoin) testnet"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "18400",
		Description: `
Audition Online Dance Battle, KAIZEN Brazil Server status/version check

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "18401",
		Description: `
Audition Online Dance Battle, KAIZEN Brazil Server

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "18505",
		Description: `
Audition Online Dance Battle R4p3 Server, Nexon Server status/version check

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "18506",
		Description: `
Audition Online Dance Battle, Nexon Server

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "18605",
		Description: `
"X-BEAT (https://en.wikipedia.org/wiki/X-BEAT) status/version check"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "18606",
		Description: `
"X-BEAT (https://en.wikipedia.org/wiki/X-BEAT)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "18676",
		Description: `
YouView

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "19000",
		Description: `
Audition Online Dance Battle, G10/alaplaya Server status/version check

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes


"JACK (https://en.wikipedia.org/wiki/JACK_Audio_Connection_Kit) sound server"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "19001",
		Description: `
Audition Online Dance Battle, G10/alaplaya Server

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "19132",
		Description: `
"Minecraft: Bedrock Edition (https://en.wikipedia.org/wiki/Minecraft:_Bedrock_Edition) multiplayer server"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "19133",
		Description: `
"Minecraft: Bedrock Edition IPv6 multiplayer server"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "19150",
		Description: `
"Gkrellm (https://en.wikipedia.org/wiki/Gkrellm) Server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "19226",
		Description: `
"Panda Software (https://en.wikipedia.org/wiki/Panda_Security) AdminSecure Communication Agent"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "19294",
		Description: `
"Google Talk (https://en.wikipedia.org/wiki/Google_Talk) Voice and Video connections"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "19295",
		Description: `
"Google Talk Voice and Video connections"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "19302",
		Description: `
"Google Talk Voice and Video connections"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "19531",
		Description: `
"systemd (https://en.wikipedia.org/wiki/Systemd)-journal-gatewayd"

IANA Status - Unofficial
TCP         - Yes
UDP         - No

`,
	},
	{
		Name: "19532",
		Description: `
"systemd (https://en.wikipedia.org/wiki/Systemd)-journal-remote"

IANA Status - Unofficial
TCP         - Yes
UDP         - No

`,
	},
	{
		Name: "19812",
		Description: `
"4D database SQL Communication"

IANA Status - Official
TCP         - Yes
UDP         - No

`,
	},
	{
		Name: "19813",
		Description: `
"4D database Client Server Communication"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "19814",
		Description: `
"4D database DB4D Communication"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "19999",
		Description: `
"Distributed Network Protocol—Secure (DNP (https://en.wikipedia.org/wiki/DNP3)—Secure), a secure version of the protocol used in SCADA (https://en.wikipedia.org/wiki/SCADA) systems between communicating RTU (https://en.wiktionary.org/wiki/RTU)'s and IED (https://en.wiktionary.org/wiki/IED)'s"

IANA Status - Official
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "20000",
		Description: `
"Distributed Network Protocol (DNP (https://en.wikipedia.org/wiki/DNP3)), a protocol used in SCADA (https://en.wikipedia.org/wiki/SCADA) systems between communicating RTU (https://en.wikipedia.org/wiki/Remote_Terminal_Unit)'s and IED (https://en.wikipedia.org/wiki/Intelligent_electronic_device)'s"

IANA Status - Official
TCP         - Unspecified
UDP         - Unspecified


"Usermin (https://en.wikipedia.org/wiki/Usermin), Web-based Unix/Linux user administration tool (default port)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified


"Used on VoIP (https://en.wikipedia.org/wiki/Voice_over_IP) networks for receiving and transmitting voice telephony traffic which includes Google Voice (https://en.wikipedia.org/wiki/Google_Voice) via the OBiTalk (https://en.wikipedia.org/wiki/Obihai_Technology) ATA (https://en.wikipedia.org/wiki/Analog_telephone_adapter) devices as well as on the MagicJack (https://en.wikipedia.org/wiki/MagicJack) and Vonage (https://en.wikipedia.org/wiki/Vonage) ATA network devices."

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "20560",
		Description: `
"Killing Floor (https://en.wikipedia.org/wiki/Killing_Floor_(2009_video_game))"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "20582",
		Description: `
HW Development IoT comms

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "20583",
		Description: `
HW Development IoT comms

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "20595",
		Description: `
"0 A.D. Empires Ascendant (https://en.wikipedia.org/wiki/0_A.D._(video_game))"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "20808",
		Description: `
Ableton Link

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "21025",
		Description: `
"Starbound Server (default), Starbound (http://playstarbound.com/)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "22000",
		Description: `
"Syncthing (https://en.wikipedia.org/wiki/Syncthing) (default)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "22136",
		Description: `
"FLIR Systems (http://www.flir.com/) Camera Resource Protocol"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "22222",
		Description: `
"Davis Instruments, WeatherLink IP (http://davisnet.com/weather/products/weather_product.asp?pnum=06555)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "23073",
		Description: `
"Soldat (https://en.wikipedia.org/wiki/Soldat_(video_game)) Dedicated Server"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "23399",
		Description: `
"Skype (https://en.wikipedia.org/wiki/Skype) default protocol"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "23513",
		Description: `
"Duke Nukem 3D source ports (/wiki/Duke_Nukem_3D#Source_ports)"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "24441",
		Description: `
Pyzor spam detection network

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "24444",
		Description: `
"NetBeans (https://en.wikipedia.org/wiki/NetBeans) integrated development environment"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "24465",
		Description: `
"Tonido Directory Server (https://en.wikipedia.org/wiki/Tonido) for Tonido (http://www.tonido.com/) which is a Personal Web App and P2P platform"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "24554",
		Description: `
"BINKP (https://en.wikipedia.org/wiki/Binkp), Fidonet (https://en.wikipedia.org/wiki/Fidonet) mail transfers over TCP/IP (https://en.wikipedia.org/wiki/TCP/IP)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "24800",
		Description: `
"Synergy (https://en.wikipedia.org/wiki/Synergy_(software)): keyboard/mouse sharing software"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "24842",
		Description: `
"StepMania: Online (https://en.wikipedia.org/wiki/StepMania): Dance Dance Revolution (https://en.wikipedia.org/wiki/Dance_Dance_Revolution) Simulator"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "25565",
		Description: `
"Minecraft (https://en.wikipedia.org/wiki/Minecraft) (Java Edition) multiplayer server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Minecraft (Java Edition) multiplayer server query"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "25575",
		Description: `
"Minecraft (Java Edition) multiplayer server RCON"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "25826",
		Description: `
"collectd (https://en.wikipedia.org/wiki/Collectd) default port"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "26000",
		Description: `
"id Software (https://en.wikipedia.org/wiki/Id_Software)'s Quake (https://en.wikipedia.org/wiki/Quake_(video_game)) server"

IANA Status - Official
TCP         - Yes
UDP         - Yes


"EVE Online (https://en.wikipedia.org/wiki/EVE_Online)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified


"Xonotic (https://en.wikipedia.org/wiki/Xonotic), an open-source (https://en.wikipedia.org/wiki/Open-source_software) arena shooter (https://en.wikipedia.org/wiki/Arena_shooter)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "26900–26901",
		Description: `
EVE Online

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "26909-26911",
		Description: `
Action Tanks Online

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "27000",
		Description: `
"PowerBuilder (https://en.wikipedia.org/wiki/PowerBuilder) SySAM (/w/index.php?title=SySAM&action=edit&redlink=1) license server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "27000–27006",
		Description: `
"id Software (https://en.wikipedia.org/wiki/Id_Software)'s QuakeWorld (https://en.wikipedia.org/wiki/QuakeWorld) master server"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "27000–27009",
		Description: `
"FlexNet Publisher (https://en.wikipedia.org/wiki/FlexNet_Publisher)'s License server (from the range of default ports)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "27000–27015",
		Description: `
"Steam (https://en.wikipedia.org/wiki/Steam_(software)) (game client traffic)"

IANA Status - Unofficial
TCP         - No
UDP         - Yes

`,
	},
	{
		Name: "27015",
		Description: `
"GoldSrc (https://en.wikipedia.org/wiki/GoldSrc) and Source engine (https://en.wikipedia.org/wiki/Source_engine) dedicated server port"

IANA Status - Unofficial
TCP         - No
UDP         - Yes

`,
	},
	{
		Name: "27015-27018",
		Description: `
"Unturned (https://en.wikipedia.org/wiki/Unturned), a survival game"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "27015–27030",
		Description: `
"Steam (matchmaking and HLTV)"

IANA Status - Unofficial
TCP         - No
UDP         - Yes


"Steam (downloads)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "27016",
		Description: `
"Magicka (https://en.wikipedia.org/wiki/Magicka) server port"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "27017",
		Description: `
"MongoDB (https://en.wikipedia.org/wiki/MongoDB) daemon process (mongod) and routing service (mongos)"

IANA Status - Unofficial
TCP         - Yes
UDP         - No

`,
	},
	{
		Name: "27031",
		Description: `
"Steam (In-Home Streaming)"

IANA Status - Unofficial
TCP         - Ports 27036 & 27037
UDP         - Yes

`,
	},
	{
		Name: "27036",
		Description: `
"Steam (In-Home Streaming)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "27037",
		Description: `
"Steam (In-Home Streaming)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Ports 27031 & 27036

`,
	},
	{
		Name: "27374",
		Description: `
"Sub7 (https://en.wikipedia.org/wiki/Sub7) default."

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "27500–27900",
		Description: `
"id Software (https://en.wikipedia.org/wiki/Id_Software)'s QuakeWorld (https://en.wikipedia.org/wiki/QuakeWorld)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "27888",
		Description: `
"Kaillera (https://en.wikipedia.org/wiki/Kaillera) server"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "27901–27910",
		Description: `
"id Software (https://en.wikipedia.org/wiki/Id_Software)'s Quake II (https://en.wikipedia.org/wiki/Quake_II) master server"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "27950",
		Description: `
"OpenArena (https://en.wikipedia.org/wiki/OpenArena) outgoing"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "27960–27969",
		Description: `
"Activision (https://en.wikipedia.org/wiki/Activision)'s Enemy Territory (https://en.wikipedia.org/wiki/Wolfenstein:_Enemy_Territory) and id Software (https://en.wikipedia.org/wiki/Id_Software)'s Quake III Arena (https://en.wikipedia.org/wiki/Quake_III_Arena), Quake III and Quake Live (https://en.wikipedia.org/wiki/Quake_Live) and some ioquake3 derived games, such as Urban Terror (OpenArena incoming)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "28001",
		Description: `
"Starsiege: Tribes (https://en.wikipedia.org/wiki/Starsiege:_Tribes)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Unspecified

`,
	},
	{
		Name: "28015",
		Description: `
"Rust (video game) (https://en.wikipedia.org/wiki/wiki/Rust_(video_game))"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "28016",
		Description: `
"Rust RCON (video game)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "28770–28771",
		Description: `
"AssaultCube Reloaded (/w/index.php?title=AssaultCube_Reloaded&action=edit&redlink=1), a video game based upon a modification of AssaultCube (https://en.wikipedia.org/wiki/AssaultCube)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "28785–28786",
		Description: `
"Cube 2: Sauerbraten (https://en.wikipedia.org/wiki/Cube_2:_Sauerbraten)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "28852",
		Description: `
"Killing Floor (https://en.wikipedia.org/wiki/Killing_Floor_(2009_video_game))"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "28910",
		Description: `
"Nintendo Wi-Fi Connection (https://en.wikipedia.org/wiki/Nintendo_Wi-Fi_Connection)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "28960",
		Description: `
"Call of Duty (https://en.wikipedia.org/wiki/Call_of_Duty); Call of Duty: United Offensive (https://en.wikipedia.org/wiki/Call_of_Duty:_United_Offensive); Call of Duty 2 (https://en.wikipedia.org/wiki/Call_of_Duty_2); Call of Duty 4: Modern Warfare (https://en.wikipedia.org/wiki/Call_of_Duty_4:_Modern_Warfare); Call of Duty: World at War (https://en.wikipedia.org/wiki/Call_of_Duty:_World_at_War) (PC platform)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "29000",
		Description: `
"Perfect World (https://en.wikipedia.org/wiki/Perfect_World_(video_game)), an adventure and fantasy MMORPG (https://en.wikipedia.org/wiki/MMORPG)"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "29070",
		Description: `
"Jedi Knight: Jedi Academy (https://en.wikipedia.org/wiki/Jedi_Knight:_Jedi_Academy) by Ravensoft (https://en.wikipedia.org/wiki/Ravensoft)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "29900–29901",
		Description: `
"Nintendo Wi-Fi Connection"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "29920",
		Description: `
"Nintendo Wi-Fi Connection"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "30000",
		Description: `
"XLink Kai P2P (https://en.wikipedia.org/wiki/XLink_Kai)"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "30033",
		Description: `
"Teamspeak3 Chat Server"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "30564",
		Description: `
"Multiplicity (https://en.wikipedia.org/wiki/Multiplicity_(software)): keyboard/mouse/clipboard sharing software"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "31337",
		Description: `
"Back Orifice (https://en.wikipedia.org/wiki/Back_Orifice) and Back Orifice 2000 (https://en.wikipedia.org/wiki/Back_Orifice_2000) remote administration tools"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "31416",
		Description: `
"BOINC (https://en.wikipedia.org/wiki/BOINC) RPC (https://en.wikipedia.org/wiki/Remote_procedure_call)"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "31438",
		Description: `
"Rocket U2 (https://en.wikipedia.org/wiki/Rocket_U2)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "31457",
		Description: `
"TetriNET (https://en.wikipedia.org/wiki/TetriNET)"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "32137",
		Description: `
"Immunet Protect (https://en.wikipedia.org/wiki/Immunet) (UDP in version 2.0, TCP since version 3.0)"

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "32400",
		Description: `
"Plex Media Server (https://en.wikipedia.org/wiki/Plex_Media_Server)"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "32764",
		Description: `
"A backdoor (https://en.wikipedia.org/wiki/Backdoor_(computing)) found on certain Linksys, Netgear and other wireless DSL modems/combination routers"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "32887",
		Description: `
"Ace of Spades (https://en.wikipedia.org/wiki/Ace_of_Spades_(video_game)), a multiplayer FPS (https://en.wikipedia.org/wiki/First-person_shooter) video game"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "32976",
		Description: `
"LogMeIn Hamachi (https://en.wikipedia.org/wiki/LogMeIn_Hamachi), a VPN (https://en.wikipedia.org/wiki/Virtual_private_network) application; also TCP port 12975 and SSL (https://en.wikipedia.org/wiki/Secure_Sockets_Layer) (TCP 443)."

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "33434",
		Description: `
"traceroute (https://en.wikipedia.org/wiki/Traceroute)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "33848",
		Description: `
"Jenkins (https://en.wikipedia.org/wiki/Jenkins_(software)), a continuous integration (https://en.wikipedia.org/wiki/Continuous_integration) (CI) tool"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "34000",
		Description: `
"Infestation: Survivor Stories (https://en.wikipedia.org/wiki/Infestation:_Survivor_Stories) (formerly known as The War Z), a multiplayer zombie video game"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "34197",
		Description: `
"Factorio (https://en.wikipedia.org/wiki/Factorio), a multiplayer survival and factory-building game"

IANA Status - Unofficial
TCP         - No
UDP         - Yes

`,
	},
	{
		Name: "35357",
		Description: `
"OpenStack Identity (https://en.wikipedia.org/wiki/OpenStack#Identity_(Keystone)) (Keystone) administration"

IANA Status - Official
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "37008",
		Description: `
"TZSP (https://en.wikipedia.org/wiki/TZSP) intrusion detection"

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes

`,
	},
	{
		Name: "40000",
		Description: `
"SafetyNET p (https://en.wikipedia.org/wiki/SafetyNET_p) – a real-time Industrial Ethernet (https://en.wikipedia.org/wiki/Industrial_Ethernet) protocol"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "41121",
		Description: `
"Tentacle Server - Pandora FMS (https://en.wikipedia.org/wiki/Pandora_FMS)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "41794",
		Description: `
"Crestron Control Port - Crestron Electronics (https://en.wikipedia.org/wiki/Crestron_Electronics)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "41795",
		Description: `
"Crestron Terminal Port - Crestron Electronics (https://en.wikipedia.org/wiki/Crestron_Electronics)"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "41796",
		Description: `
"Crestron Secure Control Port - Crestron Electronics (https://en.wikipedia.org/wiki/Crestron_Electronics)"

IANA Status - Official
TCP         - Yes
UDP         - No

`,
	},
	{
		Name: "41797",
		Description: `
"Crestron Secure Terminal Port - Crestron Electronics (https://en.wikipedia.org/wiki/Crestron_Electronics)"

IANA Status - Official
TCP         - Yes
UDP         - No

`,
	},
	{
		Name: "43110",
		Description: `
"ZeroNet (https://en.wikipedia.org/wiki/ZeroNet) web UI default port"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "43594–43595",
		Description: `
"RuneScape (https://en.wikipedia.org/wiki/RuneScape)"

IANA Status - Unofficial
TCP         - ?
UDP         - ?

`,
	},
	{
		Name: "44405",
		Description: `
"Mu Online (https://en.wikipedia.org/wiki/Mu_Online) Connect Server"

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified

`,
	},
	{
		Name: "44818",
		Description: `
"EtherNet/IP (https://en.wikipedia.org/wiki/EtherNet/IP) explicit messaging"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "47808–47823",
		Description: `
"BACnet (https://en.wikipedia.org/wiki/BACnet) Building Automation and Control Networks"

IANA Status - Official
TCP         - Yes
UDP         - Yes

`,
	},
	{
		Name: "49151",
		Description: `
"Reserved"

IANA Status - Official
TCP         - Reserved
UDP         - Reserved

`,
	},
}
