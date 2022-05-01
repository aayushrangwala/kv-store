package inmemory

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestNewStore(t *testing.T) {
	g := gomega.NewWithT(t)

	g.Expect(NewStore()).To(gomega.Equal(&inMemory{
		transactionStack: &stack{
			globalStore: make(map[string]string),
		},
	}))
}

func TestStore_Get(t *testing.T) {
	g := gomega.NewWithT(t)

	tests := []struct {
		desc              string
		activeTransaction *transaction
		globalStore       map[string]string
		expectedValue     string
		inputKey          string
	}{
		{
			desc:     "Should return not set when no latest transaction and key not available",
			inputKey: "a",
			globalStore: map[string]string{
				"b": "value2",
			},
			expectedValue: "not set",
		},
		{
			desc:     "Should return not set when key not found in active transaction",
			inputKey: "a",
			globalStore: map[string]string{
				"a": "value",
				"b": "value2",
			},
			expectedValue: "not set",
			activeTransaction: &transaction{
				localStore: map[string]string{
					"b": "value3",
				},
			},
		},
		{
			desc:     "Should return value with no active transaction",
			inputKey: "a",
			globalStore: map[string]string{
				"a": "value",
			},
			expectedValue: "value",
		},
		{
			desc:     "Should return value with active transaction",
			inputKey: "a",
			globalStore: map[string]string{
				"a": "value22",
				"b": "value33",
			},
			activeTransaction: &transaction{
				localStore: map[string]string{
					"a": "value",
				},
			},
			expectedValue: "value",
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			st := &inMemory{
				transactionStack: &stack{
					globalStore: make(map[string]string),
				},
			}

			if len(test.globalStore) != 0 {
				st.transactionStack.globalStore = test.globalStore
			}

			if test.activeTransaction != nil {
				st.transactionStack.latest = test.activeTransaction
				st.transactionStack.size++
			}

			g.Expect(st.Get(test.inputKey)).To(gomega.Equal(test.expectedValue))
		})
	}
}

func TestStore_Set(t *testing.T) {
	g := gomega.NewWithT(t)

	tests := []struct {
		desc                string
		activeTransaction   *transaction
		globalStore         map[string]string
		expectedGlobalStore map[string]string
		expectedLocalStore  map[string]string
		inputKey            string
		inputValue          string
	}{
		{
			desc:       "Should set the value in active transaction",
			inputKey:   "a",
			inputValue: "value1",
			globalStore: map[string]string{
				"a": "value2",
			},
			expectedGlobalStore: map[string]string{
				"a": "value2",
			},
			expectedLocalStore: map[string]string{
				"a": "value1",
				"b": "value3",
			},
			activeTransaction: &transaction{
				localStore: map[string]string{
					"b": "value3",
				},
			},
		},
		{
			desc:       "Should set the value in global store when no active transaction",
			inputKey:   "a",
			inputValue: "value1",
			globalStore: map[string]string{
				"a": "value2",
			},
			expectedGlobalStore: map[string]string{
				"a": "value1",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			st := &inMemory{
				transactionStack: &stack{
					globalStore: make(map[string]string),
				},
			}

			if len(test.globalStore) != 0 {
				st.transactionStack.globalStore = test.globalStore
			}

			if test.activeTransaction != nil {
				st.transactionStack.latest = test.activeTransaction
				st.transactionStack.size++
			}

			st.Set(test.inputKey, test.inputValue)

			g.Expect(st.transactionStack.globalStore).To(gomega.Equal(test.expectedGlobalStore))

			if test.activeTransaction == nil {
				g.Expect(st.transactionStack.latest).To(gomega.BeNil())

				return
			}

			g.Expect(st.transactionStack.latest.localStore).To(gomega.Equal(test.expectedLocalStore))
		})
	}
}

func TestStore_Delete(t *testing.T) {
	g := gomega.NewWithT(t)

	tests := []struct {
		desc                string
		activeTransaction   *transaction
		globalStore         map[string]string
		expectedGlobalStore map[string]string
		expectedLocalStore  map[string]string
		inputKey            string
	}{
		{
			desc:     "Should delete the value from active transaction",
			inputKey: "a",
			globalStore: map[string]string{
				"a": "value1",
			},
			expectedGlobalStore: map[string]string{
				"a": "value1",
			},
			expectedLocalStore: map[string]string{
				"b": "value3",
			},
			activeTransaction: &transaction{
				localStore: map[string]string{
					"b": "value3",
					"a": "value22",
				},
			},
		},
		{
			desc:     "Should delete the value from global store when no active transaction",
			inputKey: "a",
			globalStore: map[string]string{
				"a": "value22",
				"b": "value33",
			},
			expectedGlobalStore: map[string]string{
				"b": "value33",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			st := &inMemory{
				transactionStack: &stack{
					globalStore: make(map[string]string),
				},
			}

			if len(test.globalStore) != 0 {
				st.transactionStack.globalStore = test.globalStore
			}

			if test.activeTransaction != nil {
				st.transactionStack.latest = test.activeTransaction
				st.transactionStack.size++
			}

			st.Delete(test.inputKey)

			g.Expect(st.transactionStack.globalStore).To(gomega.Equal(test.expectedGlobalStore))

			if test.activeTransaction == nil {
				g.Expect(st.transactionStack.latest).To(gomega.BeNil())

				return
			}

			g.Expect(st.transactionStack.latest.localStore).To(gomega.Equal(test.expectedLocalStore))
		})
	}
}
