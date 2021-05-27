package main

import "testing"

func Test_allRunning(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"RUNNING", args{`[{"id": "00", "state": "RUNNING", "error": "NO_ERROR", "project": 17236, "run": 107, "clone": 87, "gen": 7, "core": "0xa7", "unit": "0x0000001680fccb02609dcd08cc773dfd", "percentdone": "24.14%", "eta": "2 hours 49 mins", "ppd": "24104", "creditestimate": "3738", "waitingon": "", "nextattempt": "0.00 secs", "timeremaining": "2.96 days", "totalframes": 100, "framesdone": 24, "assigned": "2021-05-24T14:39:19Z", "timeout": "2021-05-25T14:39:19Z", "deadline": "2021-05-27T14:39:19Z", "ws": "128.252.203.2", "cs": "0.0.0.0", "attempts": 0, "slot": "00", "tpf": "2 mins 14 secs", "basecredit": "990"}]`}, true},
		{"READY", args{`[{"id": "00", "state": "READY", "error": "NO_ERROR", "project": 17231, "run": 10613, "clone": 0, "gen": 22, "core": "0xa7", "unit": "0x00000000000000160000434f00002975", "percentdone": "0.00%", "eta": "2.00 days", "ppd": "1655", "creditestimate": "3310", "waitingon": "FahCore Run", "nextattempt": "11.92 secs", "timeremaining": "7.00 days", "totalframes": 0, "framesdone": 0, "assigned": "2021-05-24T15:36:57Z", "timeout": "2021-05-26T15:36:57Z", "deadline": "2021-05-31T15:36:57Z", "ws": "206.223.170.146", "cs": "128.252.203.10", "attempts": 1, "slot": "00", "tpf": "28 mins 48 secs", "basecredit": "3310"}]`}, false},
		{"DOWNLOAD", args{`[{"id": "00", "state": "DOWNLOAD", "error": "NO_ERROR", "project": 0, "run": 0, "clone": 0, "gen": 0, "core": "unknown", "unit": "0x00000000000000000000000000000000", "percentdone": "0.00%", "eta": "0.00 secs", "ppd": "0", "creditestimate": "0", "waitingon": "", "nextattempt": "0.00 secs", "timeremaining": "unknown time", "totalframes": 0, "framesdone": 0, "assigned": "<invalid>", "timeout": "<invalid>", "deadline": "<invalid>", "ws": "206.223.170.146", "cs": "0.0.0.0", "attempts": 0, "slot": "00", "tpf": "0.00 secs", "basecredit": "0"}]`}, false},
		{"CLEANUP", args{`[{"id": "02", "state": "CLEANUP", "error": "NO_ERROR", "project": 17236, "run": 116, "clone": 209, "gen": 5, "core": "0xa7", "unit": "0x0000000f80fccb02609dca0ead44f020", "percentdone": "100.00%", "eta": "0.00 secs", "ppd": "27901", "creditestimate": "3036", "waitingon": "Cleanup", "nextattempt": "54 mins 29 secs", "timeremaining": "2.76 days", "totalframes": 100, "framesdone": 100, "assigned": "2021-05-27T02:28:04Z", "timeout": "2021-05-28T02:28:04Z", "deadline": "2021-05-30T02:28:04Z", "ws": "128.252.203.2", "cs": "0.0.0.0", "attempts": 11, "slot": "00", "tpf": "1 mins 34 secs", "basecredit": "990"}]`}, false},
		{"RUNNING,READY", args{`[{"id": "00", "state": "RUNNING", "error": "NO_ERROR", "project": 17236, "run": 107, "clone": 87, "gen": 7, "core": "0xa7", "unit": "0x0000001680fccb02609dcd08cc773dfd", "percentdone": "24.14%", "eta": "2 hours 49 mins", "ppd": "24104", "creditestimate": "3738", "waitingon": "", "nextattempt": "0.00 secs", "timeremaining": "2.96 days", "totalframes": 100, "framesdone": 24, "assigned": "2021-05-24T14:39:19Z", "timeout": "2021-05-25T14:39:19Z", "deadline": "2021-05-27T14:39:19Z", "ws": "128.252.203.2", "cs": "0.0.0.0", "attempts": 0, "slot": "00", "tpf": "2 mins 14 secs", "basecredit": "990"}, {"id": "00", "state": "READY", "error": "NO_ERROR", "project": 17231, "run": 10613, "clone": 0, "gen": 22, "core": "0xa7", "unit": "0x00000000000000160000434f00002975", "percentdone": "0.00%", "eta": "2.00 days", "ppd": "1655", "creditestimate": "3310", "waitingon": "FahCore Run", "nextattempt": "11.92 secs", "timeremaining": "7.00 days", "totalframes": 0, "framesdone": 0, "assigned": "2021-05-24T15:36:57Z", "timeout": "2021-05-26T15:36:57Z", "deadline": "2021-05-31T15:36:57Z", "ws": "206.223.170.146", "cs": "128.252.203.10", "attempts": 1, "slot": "00", "tpf": "28 mins 48 secs", "basecredit": "3310"}]`}, false},
		{"RUNNINGx2", args{`[{"id": "00", "state": "RUNNING", "error": "NO_ERROR", "project": 17236, "run": 107, "clone": 87, "gen": 7, "core": "0xa7", "unit": "0x0000001680fccb02609dcd08cc773dfd", "percentdone": "24.14%", "eta": "2 hours 49 mins", "ppd": "24104", "creditestimate": "3738", "waitingon": "", "nextattempt": "0.00 secs", "timeremaining": "2.96 days", "totalframes": 100, "framesdone": 24, "assigned": "2021-05-24T14:39:19Z", "timeout": "2021-05-25T14:39:19Z", "deadline": "2021-05-27T14:39:19Z", "ws": "128.252.203.2", "cs": "0.0.0.0", "attempts": 0, "slot": "00", "tpf": "2 mins 14 secs", "basecredit": "990"}, {"id": "00", "state": "RUNNING", "error": "NO_ERROR", "project": 17236, "run": 107, "clone": 87, "gen": 7, "core": "0xa7", "unit": "0x0000001680fccb02609dcd08cc773dfd", "percentdone": "24.14%", "eta": "2 hours 49 mins", "ppd": "24104", "creditestimate": "3738", "waitingon": "", "nextattempt": "0.00 secs", "timeremaining": "2.96 days", "totalframes": 100, "framesdone": 24, "assigned": "2021-05-24T14:39:19Z", "timeout": "2021-05-25T14:39:19Z", "deadline": "2021-05-27T14:39:19Z", "ws": "128.252.203.2", "cs": "0.0.0.0", "attempts": 0, "slot": "00", "tpf": "2 mins 14 secs", "basecredit": "990"}]`}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := allRunning(tt.args.s); got != tt.want {
				t.Errorf("allRunning() = %v, want %v", got, tt.want)
			}
		})
	}
}
