package d30hectorda

import (
	"github.com/dminGod/HowRU/structs"
)

func GetHDDs() structs.CommandsList {

	c := structs.CommandsList{
		Commands : []structs.CommandTypes{
			{
				List : []structs.CommandMsg{
					{ Command : `ssh 127.0.0.1 "df -h" |  tail -n +2 | awk '{print $1 " " $5}'`,
						Server : "127.0.0.1" },
					},
					Command_type: "df_h",
				},
				{
					List : []structs.CommandMsg{
						{ Command : `ssh 127.0.0.1 "free -m" |  tail -n +2 | awk '{print $1 " " $5}'`,
							Server : "127.0.0.1" },
					},
				Command_type: "free_m",
			},
		},
	}

	return c
}

func ProcessResults(){

}


