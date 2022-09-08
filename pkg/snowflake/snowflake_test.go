package snowflake

import "testing"

/*
	func TestGenID(t *testing.T) {
		tests := []struct {
			name string
			want int64
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if got := GenID(); got != tt.want {
					t.Errorf("GenID() = %v, want %v", got, tt.want)
				}
			})
		}
	}
*/
func TestInit(t *testing.T) {
	type args struct {
		startTime string
		machineID int64
	}
	/*	tests := []struct {
			name    string
			args    args
			wantErr bool
		}{
			// TODO: Add test cases.
		}
		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				if err := Init(tt.args.startTime, tt.args.machineID); (err != nil) != tt.wantErr {
					t.Errorf("Init() error = %v, wantErr %v", err, tt.wantErr)
				}
			})
		}*/
	a := args{
		startTime: "2020-07-01",
		machineID: 1,
	}
	t.Run("fitst", func(t *testing.T) {
		if err := Init(a.startTime, a.machineID); err != nil {
			t.Fatal(err)
		}
		print("success")
	})
}
