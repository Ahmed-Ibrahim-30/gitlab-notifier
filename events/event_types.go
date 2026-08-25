package events

type EventType string

const (
	EventPush         EventType = "Push Hook"
	EventTag          EventType = "Tag Push Hook"
	EventMergeRequest EventType = "Merge Request Hook"
	EventRelease      EventType = "Release Hook"
	EventPipeline     EventType = "Pipeline Hook"
	EventDeployment   EventType = "Deployment Hook"
	EventJob          EventType = "Job Hook"
	EventComment      EventType = "Note Hook"
	EventIssue        EventType = "Issue Hook"
	EventMilestone    EventType = "Milestone Hook"
	EventEmoji        EventType = "Emoji Hook"
	EventFeatureFlag  EventType = "Feature Flag Hook"
	EventMember       EventType = "Member Hook"
	EventProject      EventType = "Project Hook"
	EventSubgroup     EventType = "Subgroup Hook"
)
