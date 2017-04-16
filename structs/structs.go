package structs


type CommandMsg struct {

	Command string
	Command_type string
	Server string
	Response string
	Response_state string
}

type CommandTypes struct {

	List []CommandMsg
	Command_type string
}

type CommandsList struct {

	Commands []CommandTypes
}
