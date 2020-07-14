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
			want: map[string]int64{"Рестораны": 769230800, "Аптеки": 769230800, "Такси": 769230800},
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

func TestMutexSumByMCC(t *testing.T) {
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
			want: map[string]int64{"Рестораны": 769230800, "Аптеки": 769230800, "Такси": 769230800},
		},
		{
			name: "У пользователя нет транзакций",
			args: args{MakeTransactions(1), 2},
			want: map[string]int64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MutexSumByMCC(tt.args.transactions, tt.args.userId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MutexSumByMCC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChanSumByMCC(t *testing.T) {
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
			want: map[string]int64{"Рестораны": 769230800, "Аптеки": 769230800, "Такси": 769230800},
		},
		{
			name: "У пользователя нет транзакций",
			args: args{MakeTransactions(1), 2},
			want: map[string]int64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ChanSumByMCC(tt.args.transactions, tt.args.userId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ChanSumByMCC() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMutexSumByMCC2(t *testing.T) {
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
			want: map[string]int64{"Рестораны": 769230800, "Аптеки": 769230800, "Такси": 769230800},
		},
		{
			name: "У пользователя нет транзакций",
			args: args{MakeTransactions(1), 2},
			want: map[string]int64{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MutexSumByMCC2(tt.args.transactions, tt.args.userId); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MutexSumByMCC2() = %v, want %v", got, tt.want)
			}
		})
	}
}
func BenchmarkChanSumByMCC(b *testing.B) {
	transactions := MakeTransactions(1)
	want := map[string]int64{"Рестораны": 769230800, "Аптеки": 769230800, "Такси": 769230800}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := ChanSumByMCC(transactions, 1)
		b.StopTimer()
		if !reflect.DeepEqual(result, want) {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
		b.StartTimer()
	}
}

func BenchmarkMutexSumByCategory(b *testing.B) {
	transactions := MakeTransactions(1)
	want := map[string]int64{"Рестораны": 769230800, "Аптеки": 769230800, "Такси": 769230800}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := MutexSumByMCC(transactions, 1)
		b.StopTimer()
		if !reflect.DeepEqual(result, want) {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
		b.StartTimer()
	}
}

func BenchmarkMutexSumByMCC2(b *testing.B) {
	transactions := MakeTransactions(1)
	want := map[string]int64{"Рестораны": 769230800, "Аптеки": 769230800, "Такси": 769230800}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := MutexSumByMCC2(transactions, 1)
		b.StopTimer()
		if !reflect.DeepEqual(result, want) {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
		b.StartTimer()
	}
}

func BenchmarkSumByMCC(b *testing.B) {
	transactions := MakeTransactions(1)
	want := map[string]int64{"Рестораны": 769230800, "Аптеки": 769230800, "Такси": 769230800}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := SumByMCC(transactions, 1)
		b.StopTimer()
		if !reflect.DeepEqual(result, want) {
			b.Fatalf("invalid result, got %v, want %v", result, want)
		}
		b.StartTimer()
	}
}
