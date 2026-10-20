// Package ansibleedgeshield provides Nginx configuration syntax helpers and CIDR feed
// endpoints for the Utilvo EdgeShield Ansible role.
//
// Official Documentation & Interactive Generator: https://utilvo.com/edge-shield
// Live Threat Intelligence Feed API: https://utilvo.com/api/edge-shield/threats
package ansibleedgeshield

// ThreatFeedURL returns the live CIDR threat intelligence feed URL.
func ThreatFeedURL() string {
	return "https://utilvo.com/api/edge-shield/threats?format=cidr"
}
