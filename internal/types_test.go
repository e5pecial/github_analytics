package internal

import (
	"testing"
)

func TestEvent_Unmarshal(t *testing.T) {
	type fields struct {
		ID      int64
		Type    string
		ActorId int64
		RepoId  int64
	}
	type args struct {
		row []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"correctFormat", fields{1, "a", 2, 3}, args{[]string{"1", "a", "2", "3"}}, false},
		{"notEnoughFields", fields{1, "a", 2, 3}, args{[]string{"1", "b", "2"}}, true},
		{"incorrectType", fields{1, "a", 2, 3}, args{[]string{"a", "b", "c", "d"}}, true},
		{"incorrectType2", fields{1, "a", 2, 3}, args{[]string{"1", "b", "c", "d"}}, true},
		{"incorrectType3", fields{1, "a", 2, 3}, args{[]string{"2", "b", "3", "d"}}, true},
		{"empty", fields{1, "aa", 2, 3}, args{[]string{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Event{
				ID:      tt.fields.ID,
				Type:    tt.fields.Type,
				ActorId: tt.fields.ActorId,
				RepoId:  tt.fields.RepoId,
			}
			if err := e.Unmarshal(tt.args.row); (err != nil) != tt.wantErr {
				t.Errorf("Event.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCommit_Unmarshal(t *testing.T) {
	type fields struct {
		Sha     string
		Message string
		EventId int64
	}
	type args struct {
		row []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"correctFormat", fields{"a", "b", 3}, args{[]string{"a", "b", "3"}}, false},
		{"notEnoughFields", fields{"a", "b", 3}, args{[]string{"a"}}, true},
		{"incorrectType", fields{"a", "b", 33}, args{[]string{"1", "2", "x"}}, true},
		{"empty", fields{"g", "aa", 1}, args{[]string{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Commit{
				Sha:     tt.fields.Sha,
				Message: tt.fields.Message,
				EventId: tt.fields.EventId,
			}
			if err := c.Unmarshal(tt.args.row); (err != nil) != tt.wantErr {
				t.Errorf("Commit.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRepository_Unmarshal(t *testing.T) {
	type fields struct {
		ID   int64
		Name string
	}
	type args struct {
		row []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"correctFormat", fields{1, "a"}, args{[]string{"1", "a"}}, false},
		{"notEnoughFields", fields{1, "a"}, args{[]string{"1"}}, true},
		{"incorrectType", fields{1, "a"}, args{[]string{"a", "b"}}, true},
		{"empty", fields{1, "aa"}, args{[]string{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				ID:   tt.fields.ID,
				Name: tt.fields.Name,
			}
			if err := r.Unmarshal(tt.args.row); (err != nil) != tt.wantErr {
				t.Errorf("Repository.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestActor_Unmarshal(t *testing.T) {
	type fields struct {
		ID       int64
		Username string
	}
	type args struct {
		row []string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{"correctFormat", fields{1, "a"}, args{[]string{"1", "a"}}, false},
		{"notEnoughFields", fields{1, "a"}, args{[]string{"1"}}, true},
		{"incorrectType", fields{1, "a"}, args{[]string{"a", "b"}}, true},
		{"empty", fields{1, "aa"}, args{[]string{}}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := &Actor{
				ID:       tt.fields.ID,
				Username: tt.fields.Username,
			}
			if err := a.Unmarshal(tt.args.row); (err != nil) != tt.wantErr {
				t.Errorf("Actor.Unmarshal() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
