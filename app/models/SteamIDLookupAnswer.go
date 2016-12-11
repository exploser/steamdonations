package models

type SteamIDLookupAnswer struct {
	Response struct {
		Players []struct {
			SteamID      string `json:"steamid"`
			Name         string `json:"personaname"`
			Lastlogoff   int    `json:"lastlogoff"`
			Profileurl   string `json:"profileurl"`
			Avatar       string `json:"avatar"`
			Avatarmedium string `json:"avatarmedium"`
			Avatarfull   string `json:"avatarfull"`
			Onlinestatus int    `json:"personastate"`
		} `json:"players"`
	} `json:"response"`
}

func (a SteamIDLookupAnswer) GetUser() User {
	return User{
		SteamID:      a.Response.Players[0].SteamID,
		Name:         a.Response.Players[0].Name,
		ProfileURL:   a.Response.Players[0].Profileurl,
		Avatar:       a.Response.Players[0].Avatar,
		AvatarMedium: a.Response.Players[0].Avatarmedium,
		AvatarFull:   a.Response.Players[0].Avatarfull,
	}
}
