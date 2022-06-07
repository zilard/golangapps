package sshprovisioner

import (


)

// InitUbuntuUser adds the ubuntu user if it doesn't
// already exist, updates its ~/.ssh/authorized_keys,
// and enables passwordless sudo for it.
//
// IniUbuntuUser will initially attempt to login as
// the ubuntu user, and verify the passwordless sudo
// is enabled; only if this is false will there be an
// attempt with the specified login
//
// authorizedKeys may be empty, in which case the file
// will be created and left empty.
func InitUbuntuUser(host, login, authorizedKeys string, privateKeys string,
			read io.Reader, write io.Writer) error {

	logger.Infof("initialising %q, user %q", host, login)

	// To avoid unnecessary prompting for the specified login,
	// initUbuntuUser will first attempt to ssh to the machine
	// as "ubuntu" with password authentication disabled, and
	// ensure that it can use sudo without a password
	//
	// Note that we explicitly do not allocate a PTY, so we
	// get a failure if sudo prompts.
	cmd := ssh.Command("ubuntu@"+host, []string{"sudo", "-n", "true"}, nil)
	if cmd.Run() == nil {
		logger.Infof("ubuntu user is already initialised")
		return nil
	}

	// Failed to login as ubuntu (or passwordless sudo is not enabled).
	// Use specified login, and execute the initUbuntuScript below.
	if login != "" {
		host = login + "@" + host
	}
















}



