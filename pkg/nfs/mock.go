package nfs

var Shares = []Share{
	{
		Directory: "/srv/nfs/test",
		Hosts: []Host{
			{
				Host:    "*",
				Options: []string{"rw", "sync"},
			},
		},
	},
	{
		Directory: "/srv/nfs/test2",
		Hosts: []Host{
			{
				Host:    "192.168.211.139/24",
				Options: []string{"rw", "sync", "no_root_squash"},
			},
			{
				Host:    "192.168.211.138/24",
				Options: []string{"ro"},
			},
		},
	},
}
