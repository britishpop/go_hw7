package transaction

import (
	"reflect"
	"testing"
)

func TestSumByMCC(t *testing.T) {
	type args struct {
		transactions []Transaction
		userId       int64
	}
	tests := []struct {
		name string
		args args
		want map[string]int64
	}{
		{
			name: "У пользователя есть транзакции",
			args: args{MakeTransactions(1), 1},
			want: map[string]int64{"Рестораны": 76900, "Аптеки": 77000, "Такси": 76900},
		},
		{
			name: "У пользователя нет транзакций",
			args: args{MakeTransactions(1), 2},
			want: map[string]int64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SumByMCC(tt.args.transactions, tt.args.userId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SumByMCC() = %v, want %v", got, tt.want)
			}
		})
	}
}
