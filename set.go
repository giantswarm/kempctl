package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	setCmd = &cobra.Command{
		Use:   "set <parameter> <value>",
		Short: "Set configuration values",
		Long: `Update configuration in the loadbalancer.

Some parameters:
 * version
 * admincert, sslrenegotiate, sslrenegotiate
 * OCSPPort, OCSPUseSSL, OCSPOnServerFail, OCSPServer, OCSPUrl
 * hostname, ha1hostname, ha2hostname, namserver, searchlist
 * dfltgw, dfltgwv6
 * backupday, backupenable, backuphost, backuphour, backupminute, backupminute, backuppath, backupuser
 * ntphost, time, timezone
 * irqbalance, linearesplogs, netconsole, netconsoleinterface
 * syslogcritical, syslogemergency, syslogerror, sysloginfo, syslognotice, syslogwarn
 * snmpcommunity, snmpcontact, snmpenable, snmptrapenable, snmpv1sink, snmpv2sink, snmplocation, snmpclient, snmpHATrap
 * emailcritical, emaildomain, emailemergency, emailenable, emailerror, emailinfo, emailnotice, emailpassword, emailport, emailserver, emailsslmode, emailuser, emailwarn
 * hoverhelp, motd, sessioncontrol, sessionidletime, sessionmaxfailattempts, wuidisplaylines
 * admingw, enableapi, geoclients, geosshport, sshaccess, sshiface, sshport, sshv1prot, wuiaccess, wuiiface, wuiport, geopartners, multihomedwui
 * ldapbackupserver, ldapsecurity, ldapserver, ldaprevalidateinterval, radiusbackupport, radiusbackupsecret, radiusbackupserver, radiusport, radiusrevalidateinterval, radiussecret, radiusserver, sessionlocalauth, sessionauthmode
 * addcookieport, addvia, allowemptyposts, alwayspersist, closeonerror, dropatdrainend, droponfail, expect100, rfcconform, rsarelocal, localbind, transparent, slowstart, addforwardheader, logsplitinterval
 * snat, allowupload, conntimeout, keepalive, multigw, nonlocalrs, onlydefaultroutes, resetclose, subnetorigin, subnetoriginating, tcptimestamp, routefilter
 * cachesize, hostcache, paranoia, limitinput
 * haif, hainitial, haprefered, hastyle, hatimeout, havhid, hawait, mcast, vmac, tcpfailover, cookieupdate, finalpersist, hamode, hacheck

 `,
		Run: setRun,
	}
)

func setRun(cmd *cobra.Command, args []string) {
	if len(args) < 2 {
		fmt.Fprintln(os.Stderr, "Parameter missing.")
		os.Exit(1)
	}
	if len(args) > 2 {
		fmt.Fprintln(os.Stderr, "Too many parameters.")
		os.Exit(1)
	}
	parameter := args[0]
	value := args[1]

	client := createClient()
	result, err := client.Set(parameter, value)

	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Printf("Updated %s from '%s' to '%s'\n", parameter, result, value)
	os.Exit(0)
}
