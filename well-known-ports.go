package httpref

// WellKnownPorts is the list of all known IANA reserved ports
var WellKnownPorts = References{
	{
		Name:    "Well Known Ports",
		IsTitle: true,
		Summary: "The port numbers in the range from 0 to 1023 (0 to 2^10 − 1)",
		Description: `The port numbers in the range from 0 to 1023 (0 to 210 − 1) are the well-known ports or system ports.[2] They are used by system processes that provide widely used types of network services. On Unix-like operating systems, a process must execute with superuser privileges to be able to bind a network socket to an IP address using one of the well-known ports.
https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "0",
		Description: `Reserved

IANA Status - Official
TCP         - Reserved
UDP         - Reserved
SCTP        - Unspecified


In programming APIs (not in communication between hosts), requests a system-allocated (dynamic) port

IANA Status - Unofficial
TCP         - N/A
UDP         - N/A
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "1",
		Description: `TCP Port Service Multiplexer ( https://en.wikipedia.org/wiki/TCP_Port_Service_Multiplexer ) (TCPMUX). Historic. Both TCP and UDP have been assigned to TCPMUX by IANA, but by design only TCP is specified.

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "2",
		Description: `compressnet (Management Utility)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "3",
		Description: `compressnet (Compression Process) 

IANA Status - Official
TCP - Assigned
UDP - Assigned 
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "5",
		Description: `Remote Job Entry ( https://en.wikipedia.org/wiki/Remote_Job_Entry ) was historically using socket 5 in its old socket form ( https://en.wikipedia.org/wiki/Network_socket#History ), while MIB PIM has identified it as TCP/5 and IANA has assigned both TCP and UDP 5 to it.

IANA Status - Official
TCP         - Assigned
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "7",
		Description: `Echo Protocol ( https://en.wikipedia.org/wiki/Echo_Protocol )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "9",
		Description: `Discard Protocol ( https://en.wikipedia.org/wiki/Discard_Protocol )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Yes


Wake-on-LAN ( https://en.wikipedia.org/wiki/Wake-on-LAN )

IANA Status - Unofficial
TCP         - No
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "11",
		Description: `Active Users (systat ( https://en.wikipedia.org/wiki/Systat_(protocol) ) service)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "13",
		Description: `Daytime Protocol ( https://en.wikipedia.org/wiki/Daytime_Protocol )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "15",
		Description: `Previously netstat ( https://en.wikipedia.org/wiki/Netstat ) service

IANA Status - Unofficial
TCP         - Yes
UDP         - No
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "17",
		Description: `Quote of the Day ( https://en.wikipedia.org/wiki/QOTD ) (QOTD)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "18",
		Description: `Message Send Protocol ( https://en.wikipedia.org/wiki/Message_Send_Protocol )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "19",
		Description: `Character Generator Protocol ( https://en.wikipedia.org/wiki/Character_Generator_Protocol ) (CHARGEN)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "20",
		Description: `File Transfer Protocol ( https://en.wikipedia.org/wiki/File_Transfer_Protocol ) (FTP) data transfer

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Yes

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "21",
		Description: `File Transfer Protocol ( https://en.wikipedia.org/wiki/File_Transfer_Protocol ) (FTP) control (command)

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Yes

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "22",
		Description: `Secure Shell ( https://en.wikipedia.org/wiki/Secure_Shell ) (SSH), secure logins, file transfers ( https://en.wikipedia.org/wiki/File_transfer ) (scp ( https://en.wikipedia.org/wiki/Secure_copy ), sftp ( https://en.wikipedia.org/wiki/SSH_file_transfer_protocol )) and port forwarding

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Yes

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "23",
		Description: `Telnet ( https://en.wikipedia.org/wiki/Telnet ) protocol—unencrypted text communications

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "25",
		Description: `Simple Mail Transfer Protocol ( https://en.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol ) (SMTP), used for email routing between mail servers

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "27",
		Description: `nsw-fe (NSW User System FE)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "28",
		Description: `Palo Alto Networks Panorama High Availability (HA) sync encrypted port.

IANA Status - Unofficial
TCP -  Unofficial
UDP - N/A
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "29",
		Description: `msg-icp (MSG ICP)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "31",
		Description: `msg-auth (MSG Authentication)

IANA Status - Official
TCP -  Assigned
UDP - N/A
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "33",
		Description: `dsp (Display Support Protocol)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "37",
		Description: `Time Protocol ( https://en.wikipedia.org/wiki/Time_Protocol )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "38",
		Description: `rap (Route Access Protocol)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "39",
		Description: `rlp (Resource Location Protocol)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "41",
		Description: `graphics (Graphics)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "42",
		Description: `Host Name Server Protocol ( https://en.wikipedia.org/wiki/ARPA_Host_Name_Server_Protocol )

IANA Status - Official
TCP         - Assigned
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "43",
		Description: `WHOIS ( https://en.wikipedia.org/wiki/WHOIS ) protocol

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "44",
		Description: `mpm-flags (MPM FLAGS Protocol)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "45",
		Description: `mpm (Message Processing Module [recv])

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "46",
		Description: `mpm-snd (MPM [default send])

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "47",
		Description: `Unspecified

IANA Status - Official
TCP         - Reserved
UDP         - Reserved
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "48",
		Description: `auditd (Digital Audit Daemon)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "49",
		Description: `TACACS ( https://en.wikipedia.org/wiki/TACACS ) Login Host protocol. TACACS+ ( https://en.wikipedia.org/wiki/TACACS%2B ), still in draft which is an improved but distinct version of TACACS, only uses TCP 49.

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "50",
		Description: `re-mail-ck (Remote Mail Checking Protocol)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "51",
		Description: `Historically used for Interface Message Processor ( https://en.wikipedia.org/wiki/Interface_Message_Processor ) logical address management, entry has been removed by IANA on 2013-05-25

IANA Status - Official
TCP         - Reserved
UDP         - Reserved
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "52",
		Description: `Xerox Network Systems ( https://en.wikipedia.org/wiki/Xerox_Network_Systems ) (XNS) Time Protocol. Despite this port being assigned by IANA, the service is meant to work on SPP ( https://en.wikipedia.org/wiki/Sequenced_Packet_Protocol ) (ancestor of IPX/SPX ( https://en.wikipedia.org/wiki/IPX/SPX )), instead of TCP/IP.

IANA Status - Official
TCP         - Assigned
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "53",
		Description: `Domain Name System ( https://en.wikipedia.org/wiki/Domain_Name_System ) (DNS)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "54",
		Description: `Xerox Network Systems (XNS) Clearinghouse (Name Server). Despite this port being assigned by IANA, the service is meant to work on SPP ( https://en.wikipedia.org/wiki/Sequenced_Packet_Protocol ) (ancestor of IPX/SPX ( https://en.wikipedia.org/wiki/IPX/SPX )), instead of TCP/IP.

IANA Status - Official
TCP         - Assigned
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "55",
		Description: `isi-gl (ISI Graphics Language)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "56",
		Description: `Xerox Network Systems (XNS) Authentication Protocol. Despite this port being assigned by IANA, the service is meant to work on SPP ( https://en.wikipedia.org/wiki/Sequenced_Packet_Protocol ) (ancestor of IPX/SPX ( https://en.wikipedia.org/wiki/IPX/SPX )), instead of TCP/IP.

IANA Status - Official
TCP         - Assigned
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "58",
		Description: `Xerox Network Systems (XNS) Mail. Despite this port being assigned by IANA, the service is meant to work on SPP ( https://en.wikipedia.org/wiki/Sequenced_Packet_Protocol ) (ancestor of IPX/SPX ( https://en.wikipedia.org/wiki/IPX/SPX )), instead of TCP/IP.

IANA Status - Official
TCP         - Assigned
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "61",
		Description: `Historically assigned to the NIFTP-Based Mail ( https://en.wikipedia.org/w/index.php?title=NIFTP-Based_Mail&action=edit&redlink=1 ) protocol, but was never documented in the related IEN ( https://en.wikipedia.org/wiki/Internet_Experiment_Note ). The port number entry was removed from IANA's registry on 2017-05-18.

IANA Status - Official
TCP         - Reserved
UDP         - Reserved
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "62",
		Description: `acas (ACA Services)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "63",
		Description: `whoispp (whois++)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "64",
		Description: `covia (Communications Integrator (CI))

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "65",
		Description: `tacacs-ds (TACACS-Database Service)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "66",
		Description: `sql-net (Oracle SQL*NET)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "67",
		Description: `Bootstrap Protocol ( https://en.wikipedia.org/wiki/Bootstrap_Protocol ) (BOOTP) server; also used by Dynamic Host Configuration Protocol ( https://en.wikipedia.org/wiki/Dynamic_Host_Configuration_Protocol ) (DHCP)

IANA Status - Official
TCP         - Assigned
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "68",
		Description: `Bootstrap Protocol (BOOTP) client; also used by Dynamic Host Configuration Protocol (DHCP)

IANA Status - Official
TCP         - Assigned
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "69",
		Description: `Trivial File Transfer Protocol ( https://en.wikipedia.org/wiki/Trivial_File_Transfer_Protocol ) (TFTP)

IANA Status - Official
TCP         - Assigned
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "70",
		Description: `Gopher ( https://en.wikipedia.org/wiki/Gopher_(protocol) ) protocol

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "71",
		Description: `NETRJS ( https://en.wikipedia.org/wiki/NETRJS ) protocol Ranges through 74

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "76",
		Description: `deos (Distributed External Object Store)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "78",
		Description: `vettcp (vettcp)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "79",
		Description: `Finger protocol ( https://en.wikipedia.org/wiki/Finger_protocol )

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "80",
		Description: `Hypertext Transfer Protocol ( https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol ) (HTTP)

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Yes


QUIC ( https://en.wikipedia.org/wiki/QUIC ), a transport protocol over UDP (still in draft as of April 2020[update] ( https://en.wikipedia.org/w/index.php?title=List_of_TCP_and_UDP_port_numbers&action=edit )), using stream multiplexing ( https://en.wikipedia.org/wiki/Multiplexing ), encryption by default with TLS ( https://en.wikipedia.org/wiki/Transport_Layer_Security ), and currently supporting HTTP/2 ( https://en.wikipedia.org/wiki/HTTP/2 ). QUIC has been renamed to HTTP/3 ( https://en.wikipedia.org/wiki/HTTP/3 ), which is currently an Internet Draft ( https://en.wikipedia.org/wiki/Internet_Draft ).

IANA Status - Unofficial
TCP         - No
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "81",
		Description: `TorPark ( https://en.wikipedia.org/wiki/TorPark ) onion routing ( https://en.wikipedia.org/wiki/Onion_routing )

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "82",
		Description: `TorPark ( https://en.wikipedia.org/wiki/TorPark ) control

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "83",
		Description: `MIT ML Device, networking file system 

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "84",
		Description: `ctf (Common Trace Facility)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "85",
		Description: `mit-ml-dev (MIT ML Device)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "86",
		Description: `mfcobol (Micro Focus Cobol)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "88",
		Description: `Kerberos ( https://en.wikipedia.org/wiki/Kerberos_(protocol) ) authentication system

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "89",
		Description: `su-mit-tg (SU/MIT Telnet Gateway)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "90",
		Description: `PointCast (dotcom) ( https://en.wikipedia.org/wiki/PointCast_(dotcom) )

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "91",
		Description: `mit-dov (MIT Dover Spooler)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "92",
		Description: `npp (Network Printing Protocol)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "93",
		Description: `dcp (Device Control Protocol)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "94",
		Description: `objcall (Tivoli Object Dispatcher)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "95",
		Description: `SUPDUP, terminal-independent remote login 

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "96",
		Description: `dixie (DIXIE Protocol Specification)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "97",
		Description: `swift-rvf (Swift Remote Virtual File Protocol)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "98",
		Description: `tacnews (TAC News)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "99",
		Description: `metagram (Metagram Relay)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "101",
		Description: `NIC ( https://en.wikipedia.org/wiki/History_of_the_Internet#NIC,_InterNIC,_IANA_and_ICANN ) host name ( https://en.wikipedia.org/wiki/Hostname )

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "102",
		Description: `ISO ( https://en.wikipedia.org/wiki/International_Organization_for_Standardization ) Transport Service Access Point (TSAP ( https://en.wikipedia.org/wiki/TSAP )) Class 0 protocol;

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "104",
		Description: `Digital Imaging and Communications in Medicine ( https://en.wikipedia.org/wiki/Digital_Imaging_and_Communications_in_Medicine ) (DICOM; also port 11112)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "105",
		Description: `CCSO Nameserver ( https://en.wikipedia.org/wiki/CCSO_Nameserver )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "106",
		Description: `macOS Server, (macOS) password server (https://en.wikipedia.org/wiki/MacOS_Server)

IANA Status - Official
TCP -  Unofficial
UDP - No
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "107",
		Description: `Remote User Telnet Service ( https://en.wikipedia.org/wiki/Rtelnet ) (RTelnet)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "108",
		Description: `IBM Systems Network Architecture ( https://en.wikipedia.org/wiki/Systems_Network_Architecture ) (SNA) gateway access server

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "109",
		Description: `Post Office Protocol ( https://en.wikipedia.org/wiki/Post_Office_Protocol ), version 2 (POP2)

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "110",
		Description: `Post Office Protocol, version 3 (POP3)

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "111",
		Description: `Open Network Computing Remote Procedure Call ( https://en.wikipedia.org/wiki/Open_Network_Computing_Remote_Procedure_Call ) (ONC RPC, sometimes referred to as Sun RPC)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "113",
		Description: `Ident ( https://en.wikipedia.org/wiki/Ident_protocol ), authentication service/identification protocol, used by IRC ( https://en.wikipedia.org/wiki/Internet_Relay_Chat ) servers to identify users

IANA Status - Official
TCP         - Yes
UDP         - No
SCTP        - Unspecified


Authentication Service (auth), the predecessor to identification protocol. Used to determine a user's identity of a particular TCP connection.

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "115",
		Description: `Simple File Transfer Protocol ( https://en.wikipedia.org/wiki/Simple_File_Transfer_Protocol )

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "117",
		Description: `UUCP Mapping Project ( https://en.wikipedia.org/wiki/UUCP_Mapping_Project ) (path service)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "118",
		Description: `Structured Query Language (SQL ( https://en.wikipedia.org/wiki/SQL )) Services

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "119",
		Description: `Network News Transfer Protocol ( https://en.wikipedia.org/wiki/Network_News_Transfer_Protocol ) (NNTP), retrieval of newsgroup messages

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "123",
		Description: `Network Time Protocol ( https://en.wikipedia.org/wiki/Network_Time_Protocol ) (NTP), used for time synchronization

IANA Status - Official
TCP         - Assigned
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "126",
		Description: `Formerly Unisys ( https://en.wikipedia.org/wiki/Unisys ) Unitary Login, renamed by Unisys to NXEdit. Used by Unisys Programmer's Workbench for Clearpath MCP, an IDE for Unisys MCP software development ( https://en.wikipedia.org/wiki/Unisys_MCP_programming_languages )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "135",
		Description: `DCE ( https://en.wikipedia.org/wiki/Distributed_Computing_Environment ) endpoint ( https://en.wikipedia.org/wiki/Communication_endpoint ) resolution

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified


Microsoft ( https://en.wikipedia.org/wiki/Microsoft ) EPMAP (End Point Mapper), also known as DCE/RPC ( https://en.wikipedia.org/wiki/Remote_procedure_call ) Locator service, used to remotely manage services including DHCP server ( https://en.wikipedia.org/wiki/DHCP_server ), DNS ( https://en.wikipedia.org/wiki/Domain_Name_System ) server and WINS ( https://en.wikipedia.org/wiki/Windows_Internet_Name_Service ). Also used by DCOM ( https://en.wikipedia.org/wiki/Distributed_Component_Object_Model )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "137",
		Description: `NetBIOS ( https://en.wikipedia.org/wiki/NetBIOS ) Name Service, used for name registration and resolution ( https://en.wikipedia.org/wiki/Name_resolution_(computer_systems) )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "138",
		Description: `NetBIOS Datagram Service

IANA Status - Official
TCP         - Assigned
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "139",
		Description: `NetBIOS Session Service

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "143",
		Description: `Internet Message Access Protocol ( https://en.wikipedia.org/wiki/Internet_Message_Access_Protocol ) (IMAP), management of electronic mail ( https://en.wikipedia.org/wiki/Email ) messages on a server

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "151",
		Description: `HEMS Energy Management System (https://en.wikipedia.org/wiki/Emergency_management_system)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "152",
		Description: `Background File Transfer Program ( https://en.wikipedia.org/w/index.php?title=Background_File_Transfer_Program&action=edit&redlink=1 ) (BFTP)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "153",
		Description: `Simple Gateway Monitoring Protocol ( https://en.wikipedia.org/wiki/Simple_Gateway_Monitoring_Protocol ) (SGMP), a protocol for remote inspection and alteration of gateway management information

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "156",
		Description: `Structured Query Language (SQL ( https://en.wikipedia.org/wiki/SQL )) Service

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "158",
		Description: `Distributed Mail System Protocol ( https://en.wikipedia.org/w/index.php?title=Distributed_Mail_System_Protocol&action=edit&redlink=1 ) (DMSP, sometimes referred to as Pcmail)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "161",
		Description: `Simple Network Management Protocol ( https://en.wikipedia.org/wiki/Simple_Network_Management_Protocol ) (SNMP)

IANA Status - Official
TCP         - Assigned
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "162",
		Description: `Simple Network Management Protocol ( https://en.wikipedia.org/wiki/Simple_Network_Management_Protocol ) Trap (SNMPTRAP)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "165",
		Description: `Xerox (https://en.wikipedia.org/wiki/Xerox)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "169",
		Description: `SEND (https://en.wikipedia.org/wiki/Secure_Neighbor_Discovery)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "170",
		Description: `Network PostScript ( https://en.wikipedia.org/wiki/PostScript ) print server ( https://en.wikipedia.org/wiki/Print_server )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "177",
		Description: `X Display Manager Control Protocol ( https://en.wikipedia.org/wiki/X_Display_Manager_Control_Protocol ) (XDMCP), used for remote logins to an X Display Manager ( https://en.wikipedia.org/wiki/X_display_manager_(program_type) ) server

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "179",
		Description: `Border Gateway Protocol ( https://en.wikipedia.org/wiki/Border_Gateway_Protocol ) (BGP), used to exchange routing and reachability information among autonomous systems ( https://en.wikipedia.org/wiki/Autonomous_system_(Internet) ) (AS) on the Internet ( https://en.wikipedia.org/wiki/Internet )

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Yes

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "180",
		Description: `ris (https://en.wikipedia.org/wiki/Remote_Installation_Services)

IANA Status - Official
TCP -  Assigned
UDP - Assigned
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "194",
		Description: `Internet Relay Chat ( https://en.wikipedia.org/wiki/Internet_Relay_Chat ) (IRC)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "199",
		Description: `SNMP ( https://en.wikipedia.org/wiki/SNMP ) Unix Multiplexer (SMUX)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "201",
		Description: `AppleTalk ( https://en.wikipedia.org/wiki/AppleTalk ) Routing Maintenance

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "209",
		Description: `Quick Mail Transfer Protocol ( https://en.wikipedia.org/wiki/Quick_Mail_Transfer_Protocol )

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "210",
		Description: `ANSI ( https://en.wikipedia.org/wiki/ANSI ) Z39.50 ( https://en.wikipedia.org/wiki/Z39.50 )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "213",
		Description: `Internetwork Packet Exchange ( https://en.wikipedia.org/wiki/Internetwork_Packet_Exchange ) (IPX)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "218",
		Description: `Message posting protocol (MPP)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "220",
		Description: `Internet Message Access Protocol ( https://en.wikipedia.org/wiki/Internet_Message_Access_Protocol ) (IMAP), version 3

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "259",
		Description: `Efficient Short Remote Operations (ESRO)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "262",
		Description: `Arcisdms

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "264",
		Description: `Border Gateway Multicast Protocol ( https://en.wikipedia.org/wiki/Border_Gateway_Multicast_Protocol ) (BGMP)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "280",
		Description: `http-mgmt

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "300",
		Description: `ThinLinc ( https://en.wikipedia.org/wiki/ThinLinc ) Web Access

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "308",
		Description: `Novastor Online Backup

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "311",
		Description: `Mac OS X Server ( https://en.wikipedia.org/wiki/Mac_OS_X_Server ) Admin (officially AppleShare ( https://en.wikipedia.org/wiki/AppleShare ) IP Web administration)

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "312",
		Description: `macOS Xsan administration (https://en.wikipedia.org/wiki/Xsan)

IANA Status - Unofficial
TCP -  Unofficial
UDP - No
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "318",
		Description: `PKIX Time Stamp Protocol ( https://en.wikipedia.org/wiki/Time_Stamp_Protocol ) (TSP)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "319",
		Description: `Precision Time Protocol ( https://en.wikipedia.org/wiki/Precision_Time_Protocol ) (PTP) event messages

IANA Status - Official
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "320",
		Description: `Precision Time Protocol ( https://en.wikipedia.org/wiki/Precision_Time_Protocol ) (PTP) general messages

IANA Status - Official
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "323",
		Description: `Resource Public Key Infrastructure (https://en.wikipedia.org/wiki/Resource_Public_Key_Infrastructure)

IANA Status - Official
TCP -  Yes
UDP - Yes
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "350",
		Description: `Mapping of Airline Traffic over Internet Protocol ( https://en.wikipedia.org/wiki/Mapping_of_Airline_Traffic_over_Internet_Protocol ) (MATIP) type A

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "351",
		Description: `MATIP type B

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "356",
		Description: `cloanto-net-1 (used by Cloanto Amiga Explorer and VMs)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "366",
		Description: `On-Demand Mail Relay (ODMR)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "369",
		Description: `Rpc2portmap

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "370",
		Description: `codaauth2, Coda authentication server

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified


securecast1, outgoing packets to NAI ( https://en.wikipedia.org/wiki/McAfee )'s SecureCast serversAs of 2000[update] ( https://en.wikipedia.org/w/index.php?title=List_of_TCP_and_UDP_port_numbers&action=edit )

IANA Status - Official
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "371",
		Description: `ClearCase albd

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "376",
		Description: `Amiga Envoy Network Inquiry Protocol (https://en.wikipedia.org/wiki/Amiga)

IANA Status - Official
TCP -  Yes
UDP - Yes
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "383",
		Description: `HP data alarm manager

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "384",
		Description: `A Remote Network Server System

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "387",
		Description: `AURP (AppleTalk ( https://en.wikipedia.org/wiki/AppleTalk ) Update-based Routing Protocol)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "388",
		Description: `Unidata LDM ( https://en.wikipedia.org/wiki/Local_Data_Manager ) near real-time data distribution protocol

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "389",
		Description: `Lightweight Directory Access Protocol ( https://en.wikipedia.org/wiki/Lightweight_Directory_Access_Protocol ) (LDAP)

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "399",
		Description: `Digital Equipment Corporation ( https://en.wikipedia.org/wiki/Digital_Equipment_Corporation ) DECnet+ ( https://en.wikipedia.org/w/index.php?title=DECnet%2B&action=edit&redlink=1 ) (Phase V) over TCP/IP (RFC1859)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "401",
		Description: `Uninterruptible power supply ( https://en.wikipedia.org/wiki/Uninterruptible_power_supply ) (UPS)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "427",
		Description: `Service Location Protocol ( https://en.wikipedia.org/wiki/Service_Location_Protocol ) (SLP)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "433",
		Description: `NNSP, part of Network News Transfer Protocol ( https://en.wikipedia.org/wiki/Network_News_Transfer_Protocol )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "434",
		Description: `Mobile IP ( https://en.wikipedia.org/wiki/Mobile_IP ) Agent (RFC 5944 (  https://tools.ietf.org/html/rfc5944 ))

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "443",
		Description: `Hypertext Transfer Protocol ( https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol ) over TLS/SSL ( https://en.wikipedia.org/wiki/Transport_Layer_Security ) (HTTPS ( https://en.wikipedia.org/wiki/HTTPS ))

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Yes


Quick UDP Internet Connections ( https://en.wikipedia.org/wiki/QUIC ) (QUIC), a transport protocol over UDP (still in draft as of April 2020[update] ( https://en.wikipedia.org/w/index.php?title=List_of_TCP_and_UDP_port_numbers&action=edit )), using stream multiplexing ( https://en.wikipedia.org/wiki/Multiplexing ), encryption by default with TLS ( https://en.wikipedia.org/wiki/Transport_Layer_Security ), and currently supporting HTTP/2 ( https://en.wikipedia.org/wiki/HTTP/2 ).

IANA Status - Unofficial
TCP         - No
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "444",
		Description: `Simple Network Paging Protocol ( https://en.wikipedia.org/wiki/Simple_Network_Paging_Protocol ) (SNPP), RFC 1568 (  https://tools.ietf.org/html/rfc1568 )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "445",
		Description: `Microsoft-DS (Directory Services) Active Directory ( https://en.wikipedia.org/wiki/Active_Directory ), Windows shares

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified


Microsoft-DS (Directory Services) SMB ( https://en.wikipedia.org/wiki/Server_Message_Block ) file sharing

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "464",
		Description: `Kerberos ( https://en.wikipedia.org/wiki/Kerberos_(protocol) ) Change/Set password

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "465",
		Description: `URL Rendezvous Directory for SSM (Cisco protocol)

IANA Status - Official
TCP         - Yes
UDP         - No
SCTP        - Unspecified


Authenticated SMTP ( https://en.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol ) over TLS/SSL ( https://en.wikipedia.org/wiki/Transport_Layer_Security ) (SMTPS ( https://en.wikipedia.org/wiki/SMTPS ))

IANA Status - Official
TCP         - Yes
UDP         - No
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "475",
		Description: `tcpnethaspsrv, Aladdin Knowledge Systems ( https://en.wikipedia.org/wiki/Aladdin_Knowledge_Systems ) Hasp services

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "476",
		Description: `Centro Software ERP ports. Range 476-490

IANA Status - Unofficial
TCP -  Unofficial
UDP - Unofficial
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "491",
		Description: `GO-Global remote access and application publishing software ( https://en.wikipedia.org/wiki/GO-Global )

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "497",
		Description: `Retrospect ( https://en.wikipedia.org/wiki/Retrospect_(software) )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "500",
		Description: `Internet Security Association and Key Management Protocol ( https://en.wikipedia.org/wiki/Internet_Security_Association_and_Key_Management_Protocol ) (ISAKMP) / Internet Key Exchange ( https://en.wikipedia.org/wiki/Internet_Key_Exchange ) (IKE)

IANA Status - Official
TCP         - Assigned
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "502",
		Description: `Modbus ( https://en.wikipedia.org/wiki/Modbus ) Protocol

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "504",
		Description: `Citadel ( https://en.wikipedia.org/wiki/Citadel/UX ), multiservice protocol for dedicated clients for the Citadel groupware system

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "510",
		Description: `FirstClass Protocol (FCP), used by FirstClass ( https://en.wikipedia.org/wiki/FirstClass ) client/server groupware system

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "512",
		Description: `Rexec ( https://en.wikipedia.org/wiki/Remote_Process_Execution ), Remote Process Execution

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified


comsat, together with biff ( https://en.wikipedia.org/wiki/Biff_(Unix) )

IANA Status - Official
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "513",
		Description: `rlogin ( https://en.wikipedia.org/wiki/Rlogin )

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified


Who

IANA Status - Official
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "514",
		Description: `Remote Shell ( https://en.wikipedia.org/wiki/Remote_Shell ), used to execute non-interactive commands on a remote system (Remote Shell, rsh, remsh)

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified


Syslog ( https://en.wikipedia.org/wiki/Syslog ), used for system logging

IANA Status - Official
TCP         - No
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "515",
		Description: `Line Printer Daemon ( https://en.wikipedia.org/wiki/Line_Printer_Daemon_protocol ) (LPD), print service

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "517",
		Description: `Talk

IANA Status - Official
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "518",
		Description: `NTalk

IANA Status - Official
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "520",
		Description: `efs, extended file name server

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified


Routing Information Protocol ( https://en.wikipedia.org/wiki/Routing_Information_Protocol ) (RIP)

IANA Status - Official
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "521",
		Description: `Routing Information Protocol Next Generation ( https://en.wikipedia.org/wiki/RIPng ) (RIPng)

IANA Status - Official
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "524",
		Description: `NetWare Core Protocol ( https://en.wikipedia.org/wiki/NetWare_Core_Protocol ) (NCP) is used for a variety things such as access to primary NetWare server resources, Time Synchronization, etc.

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "525",
		Description: `Timed, Timeserver ( https://en.wikipedia.org/wiki/Timeserver )

IANA Status - Official
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "530",
		Description: `Remote procedure call ( https://en.wikipedia.org/wiki/Remote_procedure_call ) (RPC)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "532",
		Description: `netnews

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "533",
		Description: `netwall, For Emergency Broadcasts

IANA Status - Official
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "540",
		Description: `Unix-to-Unix Copy Protocol (UUCP ( https://en.wikipedia.org/wiki/UUCP ))

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "542",
		Description: `commerce ( https://en.wikipedia.org/wiki/Commerce ) (Commerce Applications)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "543",
		Description: `klogin, Kerberos ( https://en.wikipedia.org/wiki/Kerberos_(protocol) ) login

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "544",
		Description: `kshell, Kerberos ( https://en.wikipedia.org/wiki/Kerberos_(protocol) ) Remote shell

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "546",
		Description: `DHCPv6 ( https://en.wikipedia.org/wiki/DHCPv6 ) client

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "547",
		Description: `DHCPv6 ( https://en.wikipedia.org/wiki/DHCPv6 ) server

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "548",
		Description: `Apple Filing Protocol ( https://en.wikipedia.org/wiki/Apple_Filing_Protocol ) (AFP) over TCP ( https://en.wikipedia.org/wiki/Transmission_Control_Protocol )

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "550",
		Description: `new-rwho, new-who

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "554",
		Description: `Real Time Streaming Protocol ( https://en.wikipedia.org/wiki/Real_Time_Streaming_Protocol ) (RTSP)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "556",
		Description: `Remotefs, RFS ( https://en.wikipedia.org/wiki/Remote_File_System ), rfs_server

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "560",
		Description: `rmonitor, Remote Monitor

IANA Status - Official
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "561",
		Description: `monitor

IANA Status - Official
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "563",
		Description: `NNTP ( https://en.wikipedia.org/wiki/NNTP ) over TLS/SSL ( https://en.wikipedia.org/wiki/Transport_Layer_Security ) (NNTPS)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "564",
		Description: `9P ( https://en.wikipedia.org/wiki/9P_(protocol) ) (Plan 9 ( https://en.wikipedia.org/wiki/Plan_9_from_Bell_Labs ))

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "585",
		Description: `Legacy use of Internet Message Access Protocol ( https://en.wikipedia.org/wiki/Internet_Message_Access_Protocol ) over TLS/SSL ( https://en.wikipedia.org/wiki/Transport_Layer_Security ) (IMAPS), now in use at port 993.

IANA Status - Unofficial
TCP         - Port 993
UDP         - ?
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "587",
		Description: `email message submission ( https://en.wikipedia.org/wiki/Mail_submission_agent ) (SMTP ( https://en.wikipedia.org/wiki/Simple_Mail_Transfer_Protocol ))

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "591",
		Description: `FileMaker ( https://en.wikipedia.org/wiki/FileMaker ) 6.0 (and later) Web Sharing (HTTP Alternate, also see port 80)

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "593",
		Description: `HTTP RPC Ep Map, Remote procedure call ( https://en.wikipedia.org/wiki/Remote_procedure_call ) over Hypertext Transfer Protocol ( https://en.wikipedia.org/wiki/Hypertext_Transfer_Protocol ), often used by Distributed Component Object Model ( https://en.wikipedia.org/wiki/Distributed_Component_Object_Model ) services and Microsoft Exchange Server ( https://en.wikipedia.org/wiki/Microsoft_Exchange_Server )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "601",
		Description: `Reliable Syslog ( https://en.wikipedia.org/wiki/Syslog ) Service — used for system logging

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "604",
		Description: `TUNNEL profile, a protocol for BEEP ( https://en.wikipedia.org/wiki/BEEP ) peers ( https://en.wikipedia.org/wiki/Peer-to-peer ) to form an application layer ( https://en.wikipedia.org/wiki/Application_layer ) tunnel ( https://en.wikipedia.org/wiki/Tunneling_protocol )

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "623",
		Description: `ASF Remote Management and Control Protocol (ASF-RMCP) & IPMI Remote Management Protocol

IANA Status - Official
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "625",
		Description: `Open Directory Proxy (ODProxy)

IANA Status - Unofficial
TCP         - Yes
UDP         - No
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "631",
		Description: `Internet Printing Protocol ( https://en.wikipedia.org/wiki/Internet_Printing_Protocol ) (IPP)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified


Common Unix Printing System ( https://en.wikipedia.org/wiki/Common_Unix_Printing_System ) (CUPS) administration console (extension to IPP)

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "635",
		Description: `RLZ DBase

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "636",
		Description: `Lightweight Directory Access Protocol ( https://en.wikipedia.org/wiki/Lightweight_Directory_Access_Protocol ) over TLS/SSL ( https://en.wikipedia.org/wiki/Transport_Layer_Security ) (LDAPS)

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "639",
		Description: `MSDP, Multicast Source Discovery Protocol ( https://en.wikipedia.org/wiki/Multicast_Source_Discovery_Protocol )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "641",
		Description: `SupportSoft Nexus Remote Command (control/listening), a proxy gateway connecting remote control traffic

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "643",
		Description: `SANity

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "646",
		Description: `Label Distribution Protocol ( https://en.wikipedia.org/wiki/Label_Distribution_Protocol ) (LDP), a routing protocol used in MPLS ( https://en.wikipedia.org/wiki/Multiprotocol_Label_Switching ) networks

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "647",
		Description: `DHCP Failover ( https://en.wikipedia.org/wiki/Dynamic_Host_Configuration_Protocol#Reliability ) protocol

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "648",
		Description: `Registry Registrar Protocol (RRP)

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "651",
		Description: `IEEE-MMS

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "653",
		Description: `SupportSoft Nexus Remote Command (data), a proxy gateway connecting remote control traffic

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "654",
		Description: `Media Management System (MMS) Media Management Protocol (MMP)

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "655",
		Description: `Tinc ( https://en.wikipedia.org/wiki/Tinc_(protocol) ) VPN daemon

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "657",
		Description: `IBM ( https://en.wikipedia.org/wiki/IBM ) RMC (Remote monitoring and Control) protocol, used by System p5 ( https://en.wikipedia.org/wiki/IBM_System_p ) AIX ( https://en.wikipedia.org/wiki/IBM_AIX ) Integrated Virtualization Manager (IVM) and Hardware Management Console ( https://en.wikipedia.org/wiki/IBM_Hardware_Management_Console ) to connect managed logical partitions (LPAR) ( https://en.wikipedia.org/wiki/LPAR ) to enable dynamic partition reconfiguration

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "660",
		Description: `Mac OS X Server ( https://en.wikipedia.org/wiki/Mac_OS_X_Server ) administration, version 10.4 and earlier

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "666",
		Description: `Doom ( https://en.wikipedia.org/wiki/Doom_(game) ), first online first-person shooter ( https://en.wikipedia.org/wiki/First-person_shooter )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified


airserv-ng (  http://www.aircrack-ng.org/doku.php?id=airserv-ng ), aircrack-ng ( https://en.wikipedia.org/wiki/Aircrack-ng )'s server for remote-controlling wireless devices

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "674",
		Description: `Application Configuration Access Protocol ( https://en.wikipedia.org/wiki/Application_Configuration_Access_Protocol ) (ACAP)

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "684",
		Description: `CORBA IIOP SSL

IANA Status - Official
TCP -  Yes
UDP - Yes
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "688",
		Description: `REALM-RUSD (ApplianceWare Server Appliance Management Protocol)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "690",
		Description: `Velneo Application Transfer Protocol (VATP)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "691",
		Description: `MS ( https://en.wikipedia.org/wiki/Microsoft ) Exchange ( https://en.wikipedia.org/wiki/Microsoft_Exchange_Server ) Routing

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "694",
		Description: `Linux-HA ( https://en.wikipedia.org/wiki/Linux-HA ) high-availability heartbeat

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "695",
		Description: `IEEE ( https://en.wikipedia.org/wiki/IEEE ) Media Management System over SSL ( https://en.wikipedia.org/wiki/Transport_Layer_Security ) (IEEE-MMS-SSL)

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "698",
		Description: `Optimized Link State Routing ( https://en.wikipedia.org/wiki/Optimized_Link_State_Routing_protocol ) (OLSR)

IANA Status - Official
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "700",
		Description: `Extensible Provisioning Protocol ( https://en.wikipedia.org/wiki/Extensible_Provisioning_Protocol ) (EPP), a protocol for communication between domain name registries ( https://en.wikipedia.org/wiki/Domain_name_registry ) and registrars ( https://en.wikipedia.org/wiki/Domain_name_registrar ) (RFC 5734 (  https://tools.ietf.org/html/rfc5734 ))

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "701",
		Description: `Link Management Protocol (LMP), a protocol that runs between a pair of nodes ( https://en.wikipedia.org/wiki/Node_(networking) ) and is used to manage traffic engineering ( https://en.wikipedia.org/wiki/Teletraffic_engineering ) (TE) links ( https://en.wikipedia.org/wiki/Telecommunications_link )

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "702",
		Description: `IRIS (Internet Registry Information Service) over BEEP ( https://en.wikipedia.org/wiki/BEEP ) (Blocks Extensible Exchange Protocol) (RFC 3983 (  https://tools.ietf.org/html/rfc3983 ))

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "706",
		Description: `Secure Internet Live Conferencing ( https://en.wikipedia.org/wiki/SILC_(protocol) ) (SILC)

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "711",
		Description: `Cisco ( https://en.wikipedia.org/wiki/Cisco ) Tag Distribution Protocol—being replaced by the MPLS Label Distribution Protocol ( https://en.wikipedia.org/wiki/Label_Distribution_Protocol )

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "712",
		Description: `Topology Broadcast based on Reverse-Path Forwarding routing protocol ( https://en.wikipedia.org/wiki/Topology_Broadcast_based_on_Reverse-Path_Forwarding_routing_protocol ) (TBRPF; RFC 3684 (  https://tools.ietf.org/html/rfc3684 ))

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "749",
		Description: `Kerberos (protocol) ( https://en.wikipedia.org/wiki/Kerberos_(protocol) ) administration

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "750",
		Description: `kerberos-iv, Kerberos ( https://en.wikipedia.org/wiki/Kerberos_(protocol) ) version IV

IANA Status - Official
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "751",
		Description: `kerberos_master, Kerberos ( https://en.wikipedia.org/wiki/Kerberos_(protocol) ) authentication

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "752",
		Description: `passwd_server, Kerberos ( https://en.wikipedia.org/wiki/Kerberos_(protocol) ) password (kpasswd) server

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "753",
		Description: `Reverse Routing Header (RRH)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified


userreg_server, Kerberos ( https://en.wikipedia.org/wiki/Kerberos_(protocol) ) userreg server

IANA Status - Unofficial
TCP         - Unspecified
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "754",
		Description: `tell send

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified


krb5_prop, Kerberos ( https://en.wikipedia.org/wiki/Kerberos_(protocol) ) v5 slave propagation

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "760",
		Description: `krbupdate [kreg], Kerberos ( https://en.wikipedia.org/wiki/Kerberos_(protocol) ) registration

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "782",
		Description: `Conserver ( https://en.wikipedia.org/wiki/Conserver ) serial-console management server

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "783",
		Description: `SpamAssassin ( https://en.wikipedia.org/wiki/SpamAssassin ) spamd daemon

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "800",
		Description: `mdbs-daemon

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "802",
		Description: `MODBUS ( https://en.wikipedia.org/wiki/Modbus )/TCP Security

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "808",
		Description: `Microsoft Net.TCP Port Sharing Service

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "829",
		Description: `Certificate Management Protocol ( https://en.wikipedia.org/wiki/Certificate_Management_Protocol )

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "830",
		Description: `NETCONF ( https://en.wikipedia.org/wiki/NETCONF ) over SSH ( https://en.wikipedia.org/wiki/Secure_Shell )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "831",
		Description: `NETCONF over BEEP ( https://en.wikipedia.org/wiki/BEEP )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "832",
		Description: `NETCONF for SOAP ( https://en.wikipedia.org/wiki/SOAP ) over HTTPS ( https://en.wikipedia.org/wiki/HTTPS )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "833",
		Description: `NETCONF for SOAP over BEEP ( https://en.wikipedia.org/wiki/BEEP )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "843",
		Description: `Adobe Flash ( https://en.wikipedia.org/wiki/Adobe_Flash )

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "847",
		Description: `DHCP Failover ( https://en.wikipedia.org/wiki/Dynamic_Host_Configuration_Protocol#Reliability ) protocol

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "848",
		Description: `Group Domain Of Interpretation (GDOI) protocol

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "853",
		Description: `DNS ( https://en.wikipedia.org/wiki/Domain_Name_System ) over TLS ( https://en.wikipedia.org/wiki/Transport_Layer_Security ) (RFC 7858 (  https://tools.ietf.org/html/rfc7858 ))

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "860",
		Description: `iSCSI ( https://en.wikipedia.org/wiki/ISCSI ) (RFC 3720 (  https://tools.ietf.org/html/rfc3720 ))

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "861",
		Description: `OWAMP control (RFC 4656 (  https://tools.ietf.org/html/rfc4656 ))

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "862",
		Description: `TWAMP control (RFC 5357 (  https://tools.ietf.org/html/rfc5357 ))

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "873",
		Description: `rsync ( https://en.wikipedia.org/wiki/Rsync ) file synchronization protocol

IANA Status - Official
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "888",
		Description: `cddbp, CD DataBase ( https://en.wikipedia.org/wiki/CD_database ) (CDDB ( https://en.wikipedia.org/wiki/CDDB )) protocol (CDDBP)

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified


IBM Endpoint Manager Remote Control

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "897",
		Description: `Brocade ( https://en.wikipedia.org/wiki/Brocade_Communications_Systems ) SMI-S RPC

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "898",
		Description: `Brocade SMI-S RPC SSL

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "902",
		Description: `VMware ESXi ( https://en.wikipedia.org/wiki/VMware_ESXi )

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "903",
		Description: `VMware ESXi

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "953",
		Description: `BIND ( https://en.wikipedia.org/wiki/BIND ) remote name daemon control (RNDC)

IANA Status - Official
TCP         - Yes
UDP         - Reserved
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "981",
		Description: `Remote HTTPS management for firewall devices running embedded Check Point VPN-1 ( https://en.wikipedia.org/wiki/Check_Point_VPN-1 ) software

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "988",
		Description: `Lustre (file system) Protocol (data). https://en.wikipedia.org/wiki/Lustre_(file_system)

IANA Status - Unofficial
TCP -  Unofficial
UDP - Unofficial
SCTP - N/A
DCCP - N/A

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "987",
		Description: `Microsoft Remote Web Workplace ( https://en.wikipedia.org/wiki/Microsoft_Remote_Web_Workplace ), a feature of Windows Small Business Server ( https://en.wikipedia.org/wiki/Windows_Small_Business_Server )

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "989",
		Description: `FTPS ( https://en.wikipedia.org/wiki/FTPS ) Protocol (data), FTP ( https://en.wikipedia.org/wiki/FTP ) over TLS/SSL ( https://en.wikipedia.org/wiki/Transport_Layer_Security )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "990",
		Description: `FTPS ( https://en.wikipedia.org/wiki/FTPS ) Protocol (control), FTP ( https://en.wikipedia.org/wiki/FTP ) over TLS/SSL ( https://en.wikipedia.org/wiki/Transport_Layer_Security )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "991",
		Description: `Netnews ( https://en.wikipedia.org/wiki/Netnews ) Administration System (NAS)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "992",
		Description: `Telnet ( https://en.wikipedia.org/wiki/Telnet ) protocol over TLS/SSL ( https://en.wikipedia.org/wiki/Transport_Layer_Security )

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "993",
		Description: `Internet Message Access Protocol ( https://en.wikipedia.org/wiki/Internet_Message_Access_Protocol ) over TLS/SSL ( https://en.wikipedia.org/wiki/Transport_Layer_Security ) (IMAPS)

IANA Status - Official
TCP         - Yes
UDP         - Assigned
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "994",
		Description: `Unspecified

IANA Status - Official
TCP         - Reserved
UDP         - Reserved
SCTP        - Unspecified


Internet Relay Chat ( https://en.wikipedia.org/wiki/Internet_Relay_Chat ) over TLS/SSL ( https://en.wikipedia.org/wiki/Transport_Layer_Security ) (IRCS). Previously assigned, but not used in common practice.

IANA Status - Unofficial
TCP         - Maybe
UDP         - Maybe
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "995",
		Description: `Post Office Protocol ( https://en.wikipedia.org/wiki/Post_Office_Protocol ) 3 over TLS/SSL ( https://en.wikipedia.org/wiki/Transport_Layer_Security ) (POP3S)

IANA Status - Official
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "1010",
		Description: `ThinLinc ( https://en.wikipedia.org/wiki/ThinLinc ) web-based administration interface

IANA Status - Unofficial
TCP         - Yes
UDP         - Unspecified
SCTP        - Unspecified

https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
	{
		Name: "1023",
		Description: `''

IANA Status - Official
TCP         - Reserved
UDP         - Reserved
SCTP        - Unspecified


z/OS ( https://en.wikipedia.org/wiki/Z/OS ) Network File System (NFS) (potentially ports 991-1023)

IANA Status - Unofficial
TCP         - Yes
UDP         - Yes
SCTP        - Unspecified
https://en.wikipedia.org/wiki/List_of_TCP_and_UDP_port_numbers#Well-known_ports`,
	},
}
