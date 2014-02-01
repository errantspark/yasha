package parser

import (
	"code.google.com/p/gogoprotobuf/proto"
	dota "github.com/dotabuff/d2rp/dota"
)

type Error string

func (e Error) Error() string { return string(e) }

type BaseEvent struct {
	Name  string
	Value int
}

type SignonPacket struct{}

func (s SignonPacket) ProtoMessage()  {}
func (s SignonPacket) Reset()         {}
func (s SignonPacket) String() string { return "" }

func (p *Parser) AsBaseEvent(commandName string) (proto.Message, error) {
	switch commandName {
	case "DEM_Stop":
		return &dota.CDemoStop{}, nil
	case "DEM_FileHeader":
		return &dota.CDemoFileHeader{}, nil
	case "DEM_FileInfo":
		return &dota.CDemoFileInfo{}, nil
	case "DEM_SyncTick":
		return &dota.CDemoSyncTick{}, nil
	case "DEM_SendTables":
		return &dota.CDemoSendTables{}, nil
	case "DEM_ClassInfo":
		return &dota.CDemoClassInfo{}, nil
	case "DEM_StringTables":
		return &dota.CDemoStringTables{}, nil
	case "DEM_Packet":
		return &dota.CDemoPacket{}, nil
	case "DEM_SignonPacket":
		return &SignonPacket{}, nil
	case "DEM_ConsoleCmd":
		return &dota.CDemoConsoleCmd{}, nil
	case "DEM_CustomData":
		return &dota.CDemoCustomData{}, nil
	case "DEM_CustomDataCallbacks":
		return &dota.CDemoCustomDataCallbacks{}, nil
	case "DEM_UserCmd":
		return &dota.CDemoUserCmd{}, nil
	case "DEM_FullPacket":
		return &dota.CDemoFullPacket{}, nil
	case "net_NOP":
		return &dota.CNETMsg_NOP{}, nil
	case "net_Disconnect":
		return &dota.CNETMsg_Disconnect{}, nil
	case "net_File":
		return &dota.CNETMsg_File{}, nil
	case "net_SplitScreenUser":
		return &dota.CNETMsg_SplitScreenUser{}, nil
	case "net_Tick":
		return &dota.CNETMsg_Tick{}, nil
	case "net_StringCmd":
		return &dota.CNETMsg_StringCmd{}, nil
	case "net_SetConVar":
		return &dota.CNETMsg_SetConVar{}, nil
	case "net_SignonState":
		return &dota.CNETMsg_SignonState{}, nil
	case "svc_ServerInfo":
		return &dota.CSVCMsg_ServerInfo{}, nil
	case "svc_SendTable":
		return &dota.CSVCMsg_SendTable{}, nil
	case "svc_ClassInfo":
		return &dota.CSVCMsg_ClassInfo{}, nil
	case "svc_SetPause":
		return &dota.CSVCMsg_SetPause{}, nil
	case "svc_CreateStringTable":
		return &dota.CSVCMsg_CreateStringTable{}, nil
	case "svc_UpdateStringTable":
		return &dota.CSVCMsg_UpdateStringTable{}, nil
	case "svc_VoiceInit":
		return &dota.CSVCMsg_VoiceInit{}, nil
	case "svc_VoiceData":
		return &dota.CSVCMsg_VoiceData{}, nil
	case "svc_Print":
		return &dota.CSVCMsg_Print{}, nil
	case "svc_Sounds":
		return &dota.CSVCMsg_Sounds{}, nil
	case "svc_SetView":
		return &dota.CSVCMsg_SetView{}, nil
	case "svc_FixAngle":
		return &dota.CSVCMsg_FixAngle{}, nil
	case "svc_CrosshairAngle":
		return &dota.CSVCMsg_CrosshairAngle{}, nil
	case "svc_BSPDecal":
		return &dota.CSVCMsg_BSPDecal{}, nil
	case "svc_SplitScreen":
		return &dota.CSVCMsg_SplitScreen{}, nil
	case "svc_UserMessage":
		return &dota.CSVCMsg_UserMessage{}, nil
	case "svc_GameEvent":
		return &dota.CSVCMsg_GameEvent{}, nil
	case "svc_PacketEntities":
		return &dota.CSVCMsg_PacketEntities{}, nil
	case "svc_TempEntities":
		return &dota.CSVCMsg_TempEntities{}, nil
	case "svc_Prefetch":
		return &dota.CSVCMsg_Prefetch{}, nil
	case "svc_Menu":
		return &dota.CSVCMsg_Menu{}, nil
	case "svc_GameEventList":
		return &dota.CSVCMsg_GameEventList{}, nil
	case "svc_GetCvarValue":
		return &dota.CSVCMsg_GetCvarValue{}, nil
	case "svc_PacketReliable":
		return &dota.CSVCMsg_PacketReliable{}, nil
	case "UM_AchievementEvent":
		return &dota.CUserMsg_AchievementEvent{}, nil
	case "UM_CloseCaption":
		return &dota.CUserMsg_CloseCaption{}, nil
	case "UM_CurrentTimescale":
		return &dota.CUserMsg_CurrentTimescale{}, nil
	case "UM_DesiredTimescale":
		return &dota.CUserMsg_DesiredTimescale{}, nil
	case "UM_Fade":
		return &dota.CUserMsg_Fade{}, nil
	case "UM_GameTitle":
		return &dota.CUserMsg_GameTitle{}, nil
	case "UM_Geiger":
		return &dota.CUserMsg_Geiger{}, nil
	case "UM_HintText":
		return &dota.CUserMsg_HintText{}, nil
	case "UM_HudMsg":
		return &dota.CUserMsg_HudMsg{}, nil
	case "UM_HudText":
		return &dota.CUserMsg_HudText{}, nil
	case "UM_KeyHintText":
		return &dota.CUserMsg_KeyHintText{}, nil
	case "UM_MessageText":
		return &dota.CUserMsg_MessageText{}, nil
	case "UM_RequestState":
		return &dota.CUserMsg_RequestState{}, nil
	case "UM_ResetHUD":
		return &dota.CUserMsg_ResetHUD{}, nil
	case "UM_Rumble":
		return &dota.CUserMsg_Rumble{}, nil
	case "UM_SayText":
		return &dota.CUserMsg_SayText{}, nil
	case "UM_SayText2":
		return &dota.CUserMsg_SayText2{}, nil
	case "UM_SayTextChannel":
		return &dota.CUserMsg_SayTextChannel{}, nil
	case "UM_Shake":
		return &dota.CUserMsg_Shake{}, nil
	case "UM_ShakeDir":
		return &dota.CUserMsg_ShakeDir{}, nil
	case "UM_StatsCrawlMsg":
		return &dota.CUserMsg_StatsCrawlMsg{}, nil
	case "UM_StatsSkipState":
		return &dota.CUserMsg_StatsSkipState{}, nil
	case "UM_TextMsg":
		return &dota.CUserMsg_TextMsg{}, nil
	case "UM_Tilt":
		return &dota.CUserMsg_Tilt{}, nil
	case "UM_Train":
		return &dota.CUserMsg_Train{}, nil
	case "UM_VGUIMenu":
		return &dota.CUserMsg_VGUIMenu{}, nil
	case "UM_VoiceMask":
		return &dota.CUserMsg_VoiceMask{}, nil
	case "UM_VoiceSubtitle":
		return &dota.CUserMsg_VoiceSubtitle{}, nil
	case "UM_SendAudio":
		return &dota.CUserMsg_SendAudio{}, nil
	case "DOTA_UM_AIDebugLine":
		return &dota.CDOTAUserMsg_AIDebugLine{}, nil
	case "DOTA_UM_ChatEvent":
		return &dota.CDOTAUserMsg_ChatEvent{}, nil
	case "DOTA_UM_CombatHeroPositions":
		return &dota.CDOTAUserMsg_CombatHeroPositions{}, nil
	case "DOTA_UM_CombatLogData":
		return &dota.CDOTAUserMsg_CombatLogData{}, nil
	case "DOTA_UM_CombatLogShowDeath":
		return &dota.CDOTAUserMsg_CombatLogShowDeath{}, nil
	case "DOTA_UM_CreateLinearProjectile":
		return &dota.CDOTAUserMsg_CreateLinearProjectile{}, nil
	case "DOTA_UM_DestroyLinearProjectile":
		return &dota.CDOTAUserMsg_DestroyLinearProjectile{}, nil
	case "DOTA_UM_DodgeTrackingProjectiles":
		return &dota.CDOTAUserMsg_DodgeTrackingProjectiles{}, nil
	case "DOTA_UM_GlobalLightColor":
		return &dota.CDOTAUserMsg_GlobalLightColor{}, nil
	case "DOTA_UM_GlobalLightDirection":
		return &dota.CDOTAUserMsg_GlobalLightDirection{}, nil
	case "DOTA_UM_InvalidCommand":
		return &dota.CDOTAUserMsg_InvalidCommand{}, nil
	case "DOTA_UM_LocationPing":
		return &dota.CDOTAUserMsg_LocationPing{}, nil
	case "DOTA_UM_MapLine":
		return &dota.CDOTAUserMsg_MapLine{}, nil
	case "DOTA_UM_MiniKillCamInfo":
		return &dota.CDOTAUserMsg_MiniKillCamInfo{}, nil
	case "DOTA_UM_MinimapDebugPoint":
		return &dota.CDOTAUserMsg_MinimapDebugPoint{}, nil
	case "DOTA_UM_MinimapEvent":
		return &dota.CDOTAUserMsg_MinimapEvent{}, nil
	case "DOTA_UM_NevermoreRequiem":
		return &dota.CDOTAUserMsg_NevermoreRequiem{}, nil
	case "DOTA_UM_OverheadEvent":
		return &dota.CDOTAUserMsg_OverheadEvent{}, nil
	case "DOTA_UM_SetNextAutobuyItem":
		return &dota.CDOTAUserMsg_SetNextAutobuyItem{}, nil
	case "DOTA_UM_SharedCooldown":
		return &dota.CDOTAUserMsg_SharedCooldown{}, nil
	case "DOTA_UM_SpectatorPlayerClick":
		return &dota.CDOTAUserMsg_SpectatorPlayerClick{}, nil
	case "DOTA_UM_TutorialTipInfo":
		return &dota.CDOTAUserMsg_TutorialTipInfo{}, nil
	case "DOTA_UM_UnitEvent":
		return &dota.CDOTAUserMsg_UnitEvent{}, nil
	case "DOTA_UM_ParticleManager":
		return &dota.CDOTAUserMsg_ParticleManager{}, nil
	case "DOTA_UM_BotChat":
		return &dota.CDOTAUserMsg_BotChat{}, nil
	case "DOTA_UM_HudError":
		return &dota.CDOTAUserMsg_HudError{}, nil
	case "DOTA_UM_ItemPurchased":
		return &dota.CDOTAUserMsg_ItemPurchased{}, nil
	case "DOTA_UM_Ping":
		return &dota.CDOTAUserMsg_Ping{}, nil
	case "DOTA_UM_ItemFound":
		return &dota.CDOTAUserMsg_ItemFound{}, nil
	case "DOTA_UM_SwapVerify":
		return &dota.CDOTAUserMsg_SwapVerify{}, nil
	case "DOTA_UM_WorldLine":
		return &dota.CDOTAUserMsg_WorldLine{}, nil
	case "DOTA_UM_TournamentDrop":
		return &dota.CDOTAUserMsg_TournamentDrop{}, nil
	case "DOTA_UM_ItemAlert":
		return &dota.CDOTAUserMsg_ItemAlert{}, nil
	case "DOTA_UM_HalloweenDrops":
		return &dota.CDOTAUserMsg_HalloweenDrops{}, nil
	case "DOTA_UM_ChatWheel":
		return &dota.CDOTAUserMsg_ChatWheel{}, nil
	case "DOTA_UM_ReceivedXmasGift":
		return &dota.CDOTAUserMsg_ReceivedXmasGift{}, nil
	case "DOTA_UM_UpdateSharedContent":
		return &dota.CDOTAUserMsg_UpdateSharedContent{}, nil
	case "DOTA_UM_TutorialRequestExp":
		return &dota.CDOTAUserMsg_TutorialRequestExp{}, nil
	case "DOTA_UM_TutorialPingMinimap":
		return &dota.CDOTAUserMsg_TutorialPingMinimap{}, nil
	case "DOTA_UM_ShowSurvey":
		return &dota.CDOTAUserMsg_ShowSurvey{}, nil
	case "DOTA_UM_TutorialFade":
		return &dota.CDOTAUserMsg_TutorialFade{}, nil
	case "DOTA_UM_AddQuestLogEntry":
		return &dota.CDOTAUserMsg_AddQuestLogEntry{}, nil
	case "DOTA_UM_SendStatPopup":
		return &dota.CDOTAUserMsg_SendStatPopup{}, nil
	case "DOTA_UM_TutorialFinish":
		return &dota.CDOTAUserMsg_TutorialFinish{}, nil
	case "DOTA_UM_SendRoshanPopup":
		return &dota.CDOTAUserMsg_SendRoshanPopup{}, nil
	case "DOTA_UM_SendGenericToolTip":
		return &dota.CDOTAUserMsg_SendGenericToolTip{}, nil
	}
	return nil, Error("Type not found: " + commandName)
}

func (p *Parser) AsBaseEventNETSVC(value int) (proto.Message, error) {
	switch value {
	case 1:
		return &dota.CNETMsg_Disconnect{}, nil
	case 2:
		return &dota.CNETMsg_File{}, nil
	case 3:
		return &dota.CNETMsg_SplitScreenUser{}, nil
	case 4:
		return &dota.CNETMsg_Tick{}, nil
	case 5:
		return &dota.CNETMsg_StringCmd{}, nil
	case 6:
		return &dota.CNETMsg_SetConVar{}, nil
	case 7:
		return &dota.CNETMsg_SignonState{}, nil
	case 8:
		return &dota.CSVCMsg_ServerInfo{}, nil
	case 9:
		return &dota.CSVCMsg_SendTable{}, nil
	case 10:
		return &dota.CSVCMsg_ClassInfo{}, nil
	case 11:
		return &dota.CSVCMsg_SetPause{}, nil
	case 12:
		return &dota.CSVCMsg_CreateStringTable{}, nil
	case 13:
		return &dota.CSVCMsg_UpdateStringTable{}, nil
	case 14:
		return &dota.CSVCMsg_VoiceInit{}, nil
	case 15:
		return &dota.CSVCMsg_VoiceData{}, nil
	case 16:
		return &dota.CSVCMsg_Print{}, nil
	case 17:
		return &dota.CSVCMsg_Sounds{}, nil
	case 18:
		return &dota.CSVCMsg_SetView{}, nil
	case 19:
		return &dota.CSVCMsg_FixAngle{}, nil
	case 20:
		return &dota.CSVCMsg_CrosshairAngle{}, nil
	case 21:
		return &dota.CSVCMsg_BSPDecal{}, nil
	case 22:
		return &dota.CSVCMsg_SplitScreen{}, nil
	case 23:
		return &dota.CSVCMsg_UserMessage{}, nil
	case 25:
		return &dota.CSVCMsg_GameEvent{}, nil
	case 26:
		return &dota.CSVCMsg_PacketEntities{}, nil
	case 27:
		return &dota.CSVCMsg_TempEntities{}, nil
	case 28:
		return &dota.CSVCMsg_Prefetch{}, nil
	case 29:
		return &dota.CSVCMsg_Menu{}, nil
	case 30:
		return &dota.CSVCMsg_GameEventList{}, nil
	case 31:
		return &dota.CSVCMsg_GetCvarValue{}, nil
	case 32:
		return &dota.CSVCMsg_PacketReliable{}, nil
	}
	return nil, Error("not found")
}

func (p *Parser) AsBaseEventBUMDUM(value int) (proto.Message, error) {
	switch value {
	case 1:
		return &dota.CUserMsg_AchievementEvent{}, nil
	case 2:
		return &dota.CUserMsg_CloseCaption{}, nil
	case 4:
		return &dota.CUserMsg_CurrentTimescale{}, nil
	case 5:
		return &dota.CUserMsg_DesiredTimescale{}, nil
	case 6:
		return &dota.CUserMsg_Fade{}, nil
	case 7:
		return &dota.CUserMsg_GameTitle{}, nil
	case 8:
		return &dota.CUserMsg_Geiger{}, nil
	case 9:
		return &dota.CUserMsg_HintText{}, nil
	case 10:
		return &dota.CUserMsg_HudMsg{}, nil
	case 11:
		return &dota.CUserMsg_HudText{}, nil
	case 12:
		return &dota.CUserMsg_KeyHintText{}, nil
	case 13:
		return &dota.CUserMsg_MessageText{}, nil
	case 14:
		return &dota.CUserMsg_RequestState{}, nil
	case 15:
		return &dota.CUserMsg_ResetHUD{}, nil
	case 16:
		return &dota.CUserMsg_Rumble{}, nil
	case 17:
		return &dota.CUserMsg_SayText{}, nil
	case 18:
		return &dota.CUserMsg_SayText2{}, nil
	case 19:
		return &dota.CUserMsg_SayTextChannel{}, nil
	case 20:
		return &dota.CUserMsg_Shake{}, nil
	case 21:
		return &dota.CUserMsg_ShakeDir{}, nil
	case 22:
		return &dota.CUserMsg_StatsCrawlMsg{}, nil
	case 23:
		return &dota.CUserMsg_StatsSkipState{}, nil
	case 24:
		return &dota.CUserMsg_TextMsg{}, nil
	case 25:
		return &dota.CUserMsg_Tilt{}, nil
	case 26:
		return &dota.CUserMsg_Train{}, nil
	case 27:
		return &dota.CUserMsg_VGUIMenu{}, nil
	case 28:
		return &dota.CUserMsg_VoiceMask{}, nil
	case 29:
		return &dota.CUserMsg_VoiceSubtitle{}, nil
	case 30:
		return &dota.CUserMsg_SendAudio{}, nil
	case 65:
		return &dota.CDOTAUserMsg_AIDebugLine{}, nil
	case 66:
		return &dota.CDOTAUserMsg_ChatEvent{}, nil
	case 67:
		return &dota.CDOTAUserMsg_CombatHeroPositions{}, nil
	case 68:
		return &dota.CDOTAUserMsg_CombatLogData{}, nil
	case 70:
		return &dota.CDOTAUserMsg_CombatLogShowDeath{}, nil
	case 71:
		return &dota.CDOTAUserMsg_CreateLinearProjectile{}, nil
	case 72:
		return &dota.CDOTAUserMsg_DestroyLinearProjectile{}, nil
	case 73:
		return &dota.CDOTAUserMsg_DodgeTrackingProjectiles{}, nil
	case 74:
		return &dota.CDOTAUserMsg_GlobalLightColor{}, nil
	case 75:
		return &dota.CDOTAUserMsg_GlobalLightDirection{}, nil
	case 76:
		return &dota.CDOTAUserMsg_InvalidCommand{}, nil
	case 77:
		return &dota.CDOTAUserMsg_LocationPing{}, nil
	case 78:
		return &dota.CDOTAUserMsg_MapLine{}, nil
	case 79:
		return &dota.CDOTAUserMsg_MiniKillCamInfo{}, nil
	case 80:
		return &dota.CDOTAUserMsg_MinimapDebugPoint{}, nil
	case 81:
		return &dota.CDOTAUserMsg_MinimapEvent{}, nil
	case 82:
		return &dota.CDOTAUserMsg_NevermoreRequiem{}, nil
	case 83:
		return &dota.CDOTAUserMsg_OverheadEvent{}, nil
	case 84:
		return &dota.CDOTAUserMsg_SetNextAutobuyItem{}, nil
	case 85:
		return &dota.CDOTAUserMsg_SharedCooldown{}, nil
	case 86:
		return &dota.CDOTAUserMsg_SpectatorPlayerClick{}, nil
	case 87:
		return &dota.CDOTAUserMsg_TutorialTipInfo{}, nil
	case 88:
		return &dota.CDOTAUserMsg_UnitEvent{}, nil
	case 89:
		return &dota.CDOTAUserMsg_ParticleManager{}, nil
	case 90:
		return &dota.CDOTAUserMsg_BotChat{}, nil
	case 91:
		return &dota.CDOTAUserMsg_HudError{}, nil
	case 92:
		return &dota.CDOTAUserMsg_ItemPurchased{}, nil
	case 93:
		return &dota.CDOTAUserMsg_Ping{}, nil
	case 94:
		return &dota.CDOTAUserMsg_ItemFound{}, nil
	case 96:
		return &dota.CDOTAUserMsg_SwapVerify{}, nil
	case 97:
		return &dota.CDOTAUserMsg_WorldLine{}, nil
	case 98:
		return &dota.CDOTAUserMsg_TournamentDrop{}, nil
	case 99:
		return &dota.CDOTAUserMsg_ItemAlert{}, nil
	case 100:
		return &dota.CDOTAUserMsg_HalloweenDrops{}, nil
	case 101:
		return &dota.CDOTAUserMsg_ChatWheel{}, nil
	case 102:
		return &dota.CDOTAUserMsg_ReceivedXmasGift{}, nil
	case 103:
		return &dota.CDOTAUserMsg_UpdateSharedContent{}, nil
	case 104:
		return &dota.CDOTAUserMsg_TutorialRequestExp{}, nil
	case 105:
		return &dota.CDOTAUserMsg_TutorialPingMinimap{}, nil
	case 107:
		return &dota.CDOTAUserMsg_ShowSurvey{}, nil
	case 108:
		return &dota.CDOTAUserMsg_TutorialFade{}, nil
	case 109:
		return &dota.CDOTAUserMsg_AddQuestLogEntry{}, nil
	case 110:
		return &dota.CDOTAUserMsg_SendStatPopup{}, nil
	case 111:
		return &dota.CDOTAUserMsg_TutorialFinish{}, nil
	case 112:
		return &dota.CDOTAUserMsg_SendRoshanPopup{}, nil
	case 113:
		return &dota.CDOTAUserMsg_SendGenericToolTip{}, nil
	}

	return nil, Error("Unknown BUMDUM")
}
