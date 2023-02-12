package packets

import "time"

type PacketData struct {
	Time                          *time.Duration
	PacketMotionData              *CarMotionData
	PacketSessionData             *PacketSessionData
	PacketLapData                 *LapData
	PacketEventData               *string
	PacketParticipantsData        *ParticipantData
	PacketCarSetupData            *CarSetupData
	PacketCarTelemetryData        *CarTelemetryData
	PacketCarStatusData           *CarStatusData
	PacketFinalClassificationData *FinalClassificationData
	PacketLobbyInfoData           *LobbyInfoData
	PacketCarDamageData           *CarDamageData
	PacketSessionHistoryData      *PacketSessionHistoryData
}
