// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"database/sql"
	"encoding/json"
	"time"

	"github.com/google/uuid"
)

type AccountSetting struct {
	AccountSettingID          string         `json:"account_setting_id"`
	Constrained               bool           `json:"constrained"`
	DataType                  string         `json:"data_type"`
	ConstrainedDefaultValue   sql.NullString `json:"constrained_default_value"`
	UnconstrainedDefaultValue sql.NullString `json:"unconstrained_default_value"`
}

type AccountSettingAllowedValue struct {
	AllowedValueID uuid.UUID `json:"allowed_value_id"`
	SettingID      int32     `json:"setting_id"`
	ItemValue      string    `json:"item_value"`
}

type AccountSettingDataType struct {
	DataTypeID string `json:"data_type_id"`
}

type AccountSettingValue struct {
	AccountSettingID   uuid.UUID      `json:"account_setting_id"`
	UserID             uuid.UUID      `json:"user_id"`
	SettingID          int32          `json:"setting_id"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	AllowedValueID     uuid.UUID      `json:"allowed_value_id"`
	UnconstrainedValue sql.NullString `json:"unconstrained_value"`
}

type AuthToken struct {
	TokenID   uuid.UUID `json:"token_id"`
	UserID    uuid.UUID `json:"user_id"`
	CreatedAt time.Time `json:"created_at"`
	ExpiresAt time.Time `json:"expires_at"`
}

type LabelColor struct {
	LabelColorID uuid.UUID `json:"label_color_id"`
	ColorHex     string    `json:"color_hex"`
	Position     float64   `json:"position"`
	Name         string    `json:"name"`
}

type Notification struct {
	NotificationID uuid.UUID       `json:"notification_id"`
	CausedBy       uuid.UUID       `json:"caused_by"`
	ActionType     string          `json:"action_type"`
	Data           json.RawMessage `json:"data"`
	CreatedOn      time.Time       `json:"created_on"`
}

type NotificationNotified struct {
	NotifiedID     uuid.UUID    `json:"notified_id"`
	NotificationID uuid.UUID    `json:"notification_id"`
	UserID         uuid.UUID    `json:"user_id"`
	Read           bool         `json:"read"`
	ReadAt         sql.NullTime `json:"read_at"`
}

type Organization struct {
	OrganizationID uuid.UUID `json:"organization_id"`
	CreatedAt      time.Time `json:"created_at"`
	Name           string    `json:"name"`
}

type PersonalProject struct {
	PersonalProjectID uuid.UUID `json:"personal_project_id"`
	ProjectID         uuid.UUID `json:"project_id"`
	UserID            uuid.UUID `json:"user_id"`
}

type Project struct {
	ProjectID uuid.UUID    `json:"project_id"`
	TeamID    uuid.UUID    `json:"team_id"`
	CreatedAt time.Time    `json:"created_at"`
	Name      string       `json:"name"`
	PublicOn  sql.NullTime `json:"public_on"`
	ShortID   string       `json:"short_id"`
}

type ProjectLabel struct {
	ProjectLabelID uuid.UUID      `json:"project_label_id"`
	ProjectID      uuid.UUID      `json:"project_id"`
	LabelColorID   uuid.UUID      `json:"label_color_id"`
	CreatedDate    time.Time      `json:"created_date"`
	Name           sql.NullString `json:"name"`
}

type ProjectMember struct {
	ProjectMemberID uuid.UUID `json:"project_member_id"`
	ProjectID       uuid.UUID `json:"project_id"`
	UserID          uuid.UUID `json:"user_id"`
	AddedAt         time.Time `json:"added_at"`
	RoleCode        string    `json:"role_code"`
}

type ProjectMemberInvited struct {
	ProjectMemberInvitedID uuid.UUID `json:"project_member_invited_id"`
	ProjectID              uuid.UUID `json:"project_id"`
	UserAccountInvitedID   uuid.UUID `json:"user_account_invited_id"`
}

type Role struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

type SystemOption struct {
	OptionID uuid.UUID      `json:"option_id"`
	Key      string         `json:"key"`
	Value    sql.NullString `json:"value"`
}

type Task struct {
	TaskID      uuid.UUID      `json:"task_id"`
	TaskGroupID uuid.UUID      `json:"task_group_id"`
	CreatedAt   time.Time      `json:"created_at"`
	Name        string         `json:"name"`
	Position    float64        `json:"position"`
	Description sql.NullString `json:"description"`
	DueDate     sql.NullTime   `json:"due_date"`
	Complete    bool           `json:"complete"`
	CompletedAt sql.NullTime   `json:"completed_at"`
	HasTime     bool           `json:"has_time"`
	ShortID     string         `json:"short_id"`
}

type TaskActivity struct {
	TaskActivityID uuid.UUID       `json:"task_activity_id"`
	Active         bool            `json:"active"`
	TaskID         uuid.UUID       `json:"task_id"`
	CreatedAt      time.Time       `json:"created_at"`
	CausedBy       uuid.UUID       `json:"caused_by"`
	ActivityTypeID int32           `json:"activity_type_id"`
	Data           json.RawMessage `json:"data"`
}

type TaskActivityType struct {
	TaskActivityTypeID int32  `json:"task_activity_type_id"`
	Code               string `json:"code"`
	Template           string `json:"template"`
}

type TaskAssigned struct {
	TaskAssignedID uuid.UUID `json:"task_assigned_id"`
	TaskID         uuid.UUID `json:"task_id"`
	UserID         uuid.UUID `json:"user_id"`
	AssignedDate   time.Time `json:"assigned_date"`
}

type TaskChecklist struct {
	TaskChecklistID uuid.UUID `json:"task_checklist_id"`
	TaskID          uuid.UUID `json:"task_id"`
	CreatedAt       time.Time `json:"created_at"`
	Name            string    `json:"name"`
	Position        float64   `json:"position"`
}

type TaskChecklistItem struct {
	TaskChecklistItemID uuid.UUID    `json:"task_checklist_item_id"`
	TaskChecklistID     uuid.UUID    `json:"task_checklist_id"`
	CreatedAt           time.Time    `json:"created_at"`
	Complete            bool         `json:"complete"`
	Name                string       `json:"name"`
	Position            float64      `json:"position"`
	DueDate             sql.NullTime `json:"due_date"`
}

type TaskComment struct {
	TaskCommentID uuid.UUID    `json:"task_comment_id"`
	TaskID        uuid.UUID    `json:"task_id"`
	CreatedAt     time.Time    `json:"created_at"`
	UpdatedAt     sql.NullTime `json:"updated_at"`
	CreatedBy     uuid.UUID    `json:"created_by"`
	Pinned        bool         `json:"pinned"`
	Message       string       `json:"message"`
}

type TaskDueDateReminder struct {
	DueDateReminderID uuid.UUID `json:"due_date_reminder_id"`
	TaskID            uuid.UUID `json:"task_id"`
	Period            int32     `json:"period"`
	Duration          string    `json:"duration"`
}

type TaskDueDateReminderDuration struct {
	Code string `json:"code"`
}

type TaskGroup struct {
	TaskGroupID uuid.UUID `json:"task_group_id"`
	ProjectID   uuid.UUID `json:"project_id"`
	CreatedAt   time.Time `json:"created_at"`
	Name        string    `json:"name"`
	Position    float64   `json:"position"`
}

type TaskLabel struct {
	TaskLabelID    uuid.UUID `json:"task_label_id"`
	TaskID         uuid.UUID `json:"task_id"`
	ProjectLabelID uuid.UUID `json:"project_label_id"`
	AssignedDate   time.Time `json:"assigned_date"`
}

type TaskWatcher struct {
	TaskWatcherID uuid.UUID `json:"task_watcher_id"`
	TaskID        uuid.UUID `json:"task_id"`
	UserID        uuid.UUID `json:"user_id"`
	WatchedAt     time.Time `json:"watched_at"`
}

type Team struct {
	TeamID         uuid.UUID `json:"team_id"`
	CreatedAt      time.Time `json:"created_at"`
	Name           string    `json:"name"`
	OrganizationID uuid.UUID `json:"organization_id"`
}

type TeamMember struct {
	TeamMemberID uuid.UUID `json:"team_member_id"`
	TeamID       uuid.UUID `json:"team_id"`
	UserID       uuid.UUID `json:"user_id"`
	Addeddate    time.Time `json:"addeddate"`
	RoleCode     string    `json:"role_code"`
}

type UserAccount struct {
	UserID           uuid.UUID      `json:"user_id"`
	CreatedAt        time.Time      `json:"created_at"`
	Email            string         `json:"email"`
	Username         string         `json:"username"`
	PasswordHash     string         `json:"password_hash"`
	ProfileBgColor   string         `json:"profile_bg_color"`
	FullName         string         `json:"full_name"`
	Initials         string         `json:"initials"`
	ProfileAvatarUrl sql.NullString `json:"profile_avatar_url"`
	RoleCode         string         `json:"role_code"`
	Bio              string         `json:"bio"`
	Active           bool           `json:"active"`
}

type UserAccountConfirmToken struct {
	ConfirmTokenID uuid.UUID `json:"confirm_token_id"`
	Email          string    `json:"email"`
}

type UserAccountInvited struct {
	UserAccountInvitedID uuid.UUID `json:"user_account_invited_id"`
	Email                string    `json:"email"`
	InvitedOn            time.Time `json:"invited_on"`
	HasJoined            bool      `json:"has_joined"`
}
