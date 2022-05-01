package inmemory

import (
	"testing"

	"github.com/onsi/gomega"
)

func TestStack_Start(t *testing.T) {
	g := gomega.NewWithT(t)

	tests := []struct {
		desc                string
		latest              *transaction
		expectedTransaction *transaction
		expectedSize        int32
	}{
		{
			desc: "Should start a new NESTED transaction",
			latest: &transaction{
				localStore: map[string]string{
					"b": "value3",
					"a": "value22",
				},
			},
			expectedTransaction: &transaction{
				localStore: map[string]string{},
				parent: &transaction{
					localStore: map[string]string{
						"b": "value3",
						"a": "value22",
					},
				},
			},
			expectedSize: 2,
		},
		{
			desc: "Should start a new and first transaction",
			expectedTransaction: &transaction{
				localStore: map[string]string{},
			},
			expectedSize: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			st := &stack{}

			if test.latest != nil {
				st.latest = test.latest
				st.size++
			}

			st.Start()

			g.Expect(st.latest).To(gomega.Equal(test.expectedTransaction))
			g.Expect(st.size).To(gomega.Equal(test.expectedSize))
		})
	}
}

func TestStack_Commit(t *testing.T) {
	g := gomega.NewWithT(t)

	tests := []struct {
		desc                string
		latest              *transaction
		expectedTransaction *transaction
		globalStore         map[string]string
		expectedGlobalStore map[string]string
	}{
		{
			desc: "Should return with no update when no active transaction",
			globalStore: map[string]string{
				"a": "value1",
				"b": "value2",
			},
			expectedGlobalStore: map[string]string{
				"a": "value1",
				"b": "value2",
			},
		},
		{
			desc: "Should just return as no key values to commit",
			globalStore: map[string]string{
				"a": "value1",
				"b": "value2",
			},
			expectedGlobalStore: map[string]string{
				"a": "value1",
				"b": "value2",
			},
			latest: &transaction{
				localStore: map[string]string{},
			},
			expectedTransaction: &transaction{
				localStore: map[string]string{},
			},
		},
		{
			desc: "Should update the latest and parent transactions",
			globalStore: map[string]string{
				"a": "value1",
				"b": "value2",
			},
			expectedGlobalStore: map[string]string{
				"a": "value33",
				"b": "value44",
				"c": "value55",
			},
			latest: &transaction{
				localStore: map[string]string{
					"a": "value33",
					"b": "value44",
					"c": "value55",
				},
				parent: &transaction{
					localStore: map[string]string{
						"a": "value13",
						"b": "value24",
						"c": "value35",
					},
				},
			},
			expectedTransaction: &transaction{
				localStore: map[string]string{
					"a": "value33",
					"b": "value44",
					"c": "value55",
				},
				parent: &transaction{
					localStore: map[string]string{
						"a": "value33",
						"b": "value44",
						"c": "value55",
					},
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			st := &stack{}

			if len(test.globalStore) > 0 {
				st.globalStore = test.globalStore
			}

			if test.latest != nil {
				st.latest = test.latest
				st.size++
			}

			st.Commit()

			g.Expect(st.globalStore).To(gomega.Equal(test.expectedGlobalStore))
			g.Expect(st.latest).To(gomega.Equal(test.expectedTransaction))
		})
	}
}

func TestStack_Abort(t *testing.T) {
	g := gomega.NewWithT(t)

	tests := []struct {
		desc                string
		latest              *transaction
		expectedTransaction *transaction
	}{
		{
			desc: "Should return with no abort when no active transaction",
		},
		{
			desc: "Should abort the current transaction and latest transaction will be nil when no active parent transaction",
			latest: &transaction{
				localStore: map[string]string{
					"a": "value33",
					"b": "value44",
					"c": "value55",
				},
			},
		},
		{
			desc: "Should abort the current transaction and parent transaction will be active",
			latest: &transaction{
				localStore: map[string]string{
					"a": "value33",
					"b": "value44",
					"c": "value55",
				},
				parent: &transaction{
					localStore: map[string]string{
						"a": "value13",
						"b": "value24",
						"c": "value35",
					},
				},
			},
			expectedTransaction: &transaction{
				localStore: map[string]string{
					"a": "value13",
					"b": "value24",
					"c": "value35",
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			st := &stack{}

			if test.latest != nil {
				st.latest = test.latest
				st.size++
			}

			st.Abort()

			if test.expectedTransaction == nil {
				g.Expect(st.latest).To(gomega.BeNil())

				return
			}

			g.Expect(st.latest).To(gomega.Equal(test.expectedTransaction))
		})
	}
}
