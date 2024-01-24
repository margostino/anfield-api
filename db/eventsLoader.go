package db

import (
	"github.com/margostino/anfield-api/fpl"
)

func loadEvents(fplEvents []*fpl.Event, fplFixturesResponse fpl.FixturesResponse, teamIndex map[int]*Team) map[int]*Event {
	events := make(map[int]*Event)
	for _, event := range fplEvents {
		chipPlays := make([]*ChipPlay, len(event.ChipPlays))
		for _, fplChipPlay := range event.ChipPlays {
			chipPlays = append(chipPlays, &ChipPlay{
				ChipName:  fplChipPlay.ChipName,
				NumPlayed: fplChipPlay.NumPlayed,
			})
		}

		var topElementInfo *TopElementInfo
		if event.TopElementInfo != nil {
			topElementInfo = &TopElementInfo{
				ID:     event.TopElementInfo.ID,
				Points: event.TopElementInfo.Points,
			}
		} else {
			topElementInfo = nil
		}

		fixtures := loadFixtures(fplFixturesResponse, teamIndex)

		eventFixtures := fixtures[event.ID]
		events[event.ID] = &Event{
			ID:                     event.ID,
			Name:                   event.Name,
			DeadlineTime:           event.DeadlineTime.String(),
			AverageEntryScore:      event.AverageEntryScore,
			Finished:               event.Finished,
			DataChecked:            event.DataChecked,
			HighestScoringEntry:    event.HighestScoringEntry,
			DeadlineTimeEpoch:      event.DeadlineTimeEpoch,
			DeadlineTimeGameOffset: event.DeadlineTimeGameOffset,
			HighestScore:           event.HighestScore,
			IsPrevious:             event.IsPrevious,
			IsCurrent:              event.IsCurrent,
			IsNext:                 event.IsNext,
			CupLeaguesCreated:      event.CupLeaguesCreated,
			H2hKoMatchesCreated:    event.H2hKoMatchesCreated,
			RankedCount:            event.RankedCount,
			ChipPlays:              chipPlays,
			MostSelected:           event.MostSelected,
			MostTransferredIn:      event.MostTransferredIn,
			TopElement:             event.TopElement,
			TopElementInfo:         topElementInfo,
			TransfersMade:          event.TransfersMade,
			MostCaptained:          event.MostCaptained,
			MostViceCaptained:      event.MostViceCaptained,
			Fixtures:               eventFixtures,
		}
	}
	return events
}
