package tracker

import (
	"github.com/topi314/campfire-tools/server/database"
)

func categoryFilterOptions() []string {
	options := make([]string, 0, len(orderedEventCategories)+2)
	options = append(options, orderedEventCategories...)
	return append(options, EventCategoryOther, EventCategoryNoEvent)
}

func (h *handler) filterEventsByCategory(events []database.EventWithCheckIns, category string) []database.EventWithCheckIns {
	if category == "" {
		return events
	}

	filtered := make([]database.EventWithCheckIns, 0, len(events))
	for _, event := range events {
		if h.getEventCategory(event.CampfireLiveEventName) == category {
			filtered = append(filtered, event)
		}
	}
	return filtered
}
