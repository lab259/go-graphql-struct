package gqlstruct_test

import (
	"github.com/graphql-go/graphql"
	"github.com/lab259/go-graphql-struct"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"time"
)

type CustomFieldType struct {
	Value string
}

func (t *CustomFieldType) GraphqlType() graphql.Type {
	return graphql.Float
}

type CustomFieldTypeWithResolver struct {
	Value string
}

func (t *CustomFieldTypeWithResolver) GraphqlResolve(p graphql.ResolveParams) (interface{}, error) {
	panic("only to catch")
}

func (t *CustomFieldTypeWithResolver) GraphqlType() graphql.Type {
	return graphql.Float
}

var _ = Describe("Struct", func() {
	It("should ignore not tagged fields", func() {
		type StructExample struct {
			Field1 string `graphql:"field1"`
			Field2 int
			Field3 bool `graphql:"field3"`
		}

		obj := gqlstruct.Struct(&StructExample{})
		Expect(obj).ToNot(BeNil())
		Expect(obj.Name()).To(Equal("StructExample"))
		fields := obj.Fields()
		Expect(fields).ToNot(BeNil())
		Expect(fields).To(HaveLen(2))
		Expect(fields).To(HaveKey("field1"))
		Expect(fields).To(HaveKey("field3"))
		Expect(fields["field1"].Name).To(Equal("field1"))
		Expect(fields["field1"].Type.String()).To(Equal("String"))
		Expect(fields["field3"].Name).To(Equal("field3"))
		Expect(fields["field3"].Type.String()).To(Equal("Boolean"))
	})

	Context("Custom Type", func() {
		It("should generate graphql.Type with a custom type field", func() {
			type StructExample struct {
				Field1 *CustomFieldType `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Float"))
		})

		It("should generate graphql.Type with a non null custom type field", func() {
			type StructExample struct {
				Field1 CustomFieldType `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Float!"))
		})
	})

	Context("String Type", func() {

		It("should generate graphql.Type with a string field", func() {
			type StructExample struct {
				Field1 *string `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("String"))
		})

		It("should generate graphql.Type with a non null string field", func() {
			type StructExample struct {
				Field1 string `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("String!"))
		})
	})

	Context("Int Type", func() {
		It("should generate graphql.Type with a int field", func() {
			type StructExample struct {
				Field1 *int `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int"))
		})

		It("should generate graphql.Type with a non null int field", func() {
			type StructExample struct {
				Field1 int `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int!"))
		})

		It("should generate graphql.Type with a int8 field", func() {
			type StructExample struct {
				Field1 *int8 `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int"))
		})

		It("should generate graphql.Type with a non null int8 field", func() {
			type StructExample struct {
				Field1 int8 `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int!"))
		})

		It("should generate graphql.Type with a int16 field", func() {
			type StructExample struct {
				Field1 *int16 `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int"))
		})

		It("should generate graphql.Type with a non null int16 field", func() {
			type StructExample struct {
				Field1 int16 `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int!"))
		})

		It("should generate graphql.Type with a int32 field", func() {
			type StructExample struct {
				Field1 *int32 `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int"))
		})

		It("should generate graphql.Type with a non null int32 field", func() {
			type StructExample struct {
				Field1 int32 `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int!"))
		})

		It("should generate graphql.Type with a int64 field", func() {
			type StructExample struct {
				Field1 *int64 `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int"))
		})

		It("should generate graphql.Type with a non null int64 field", func() {
			type StructExample struct {
				Field1 int64 `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int!"))
		})

		It("should generate graphql.Type with a uint field", func() {
			type StructExample struct {
				Field1 *uint `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int"))
		})

		It("should generate graphql.Type with a non null uint field", func() {
			type StructExample struct {
				Field1 uint `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int!"))
		})

		It("should generate graphql.Type with a uint8 field", func() {
			type StructExample struct {
				Field1 *uint8 `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int"))
		})

		It("should generate graphql.Type with a non null uint8 field", func() {
			type StructExample struct {
				Field1 uint8 `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int!"))
		})

		It("should generate graphql.Type with a uint16 field", func() {
			type StructExample struct {
				Field1 *uint16 `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int"))
		})

		It("should generate graphql.Type with a non null uint16 field", func() {
			type StructExample struct {
				Field1 uint16 `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int!"))
		})

		It("should generate graphql.Type with a uint32 field", func() {
			type StructExample struct {
				Field1 *uint32 `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int"))
		})

		It("should generate graphql.Type with a non null uint32 field", func() {
			type StructExample struct {
				Field1 uint32 `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int!"))
		})

		It("should generate graphql.Type with a uint64 field", func() {
			type StructExample struct {
				Field1 *uint64 `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int"))
		})

		It("should generate graphql.Type with a non null uint64 field", func() {
			type StructExample struct {
				Field1 uint64 `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Int!"))
		})
	})

	Context("Float Type", func() {
		It("should generate graphql.Type with a float32 field", func() {
			type StructExample struct {
				Field1 *float32 `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Float"))
		})

		It("should generate graphql.Type with a non null float32 field", func() {
			type StructExample struct {
				Field1 float32 `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Float!"))
		})

		It("should generate graphql.Type with a float64 field", func() {
			type StructExample struct {
				Field1 *float64 `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Float"))
		})

		It("should generate graphql.Type with a non null float64 field", func() {
			type StructExample struct {
				Field1 float64 `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Float!"))
		})

		It("should generate graphql.Type with a complex64 field", func() {
			type StructExample struct {
				Field1 *complex64 `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Float"))
		})

		It("should generate graphql.Type with a non null complex64 field", func() {
			type StructExample struct {
				Field1 complex64 `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Float!"))
		})

		It("should generate graphql.Type with a complex128 field", func() {
			type StructExample struct {
				Field1 *complex128 `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Float"))
		})

		It("should generate graphql.Type with a non null complex128 field", func() {
			type StructExample struct {
				Field1 complex128 `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Float!"))
		})
	})

	Context("Boolean Type", func() {
		It("should generate graphql.Type with a bool field", func() {
			type StructExample struct {
				Field1 *bool `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Boolean"))
		})

		It("should generate graphql.Type with a non null bool field", func() {
			type StructExample struct {
				Field1 bool `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("Boolean!"))
		})
	})

	Context("Time Type", func() {
		It("should generate graphql.Type with a time.Time field", func() {
			type StructExample struct {
				Field1 *time.Time `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("DateTime"))
		})

		It("should generate graphql.Type with a non null time.Time field", func() {
			type StructExample struct {
				Field1 time.Time `graphql:"!field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj).ToNot(BeNil())
			Expect(obj.Name()).To(Equal("StructExample"))
			fields := obj.Fields()
			Expect(fields).ToNot(BeNil())
			Expect(fields).To(HaveLen(1))
			Expect(fields).To(HaveKey("field1"))
			Expect(fields["field1"].Name).To(Equal("field1"))
			Expect(fields["field1"].Type.String()).To(Equal("DateTime!"))
		})
	})

	Context("Unknown Type", func() {
		type CustomFieldTypeWithNoGraphQL struct {
			Value string
		}

		It("should panic with a pointer to a unknown type", func() {
			type StructExample struct {
				Field1 *CustomFieldTypeWithNoGraphQL `graphql:"field1"`
			}

			Expect(func() {
				gqlstruct.Struct(&StructExample{})
			}).To(Panic())
		})

		It("should panic with a unknown type", func() {
			type StructExample struct {
				Field1 CustomFieldTypeWithNoGraphQL `graphql:"field1"`
			}

			Expect(func() {
				gqlstruct.Struct(&StructExample{})
			}).To(Panic())
		})
	})

	Context("Resolvers", func() {
		It("should set the resolver of a pointer field", func() {
			type StructExample struct {
				Field1 *CustomFieldTypeWithResolver `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj.Fields()).To(HaveLen(1))
			Expect(obj.Fields()).To(HaveKey("field1"))
			func() {
				defer func() {
					r := recover()
					Expect(r).To(Equal("only to catch"))
				}()

				obj.Fields()["field1"].Resolve(graphql.ResolveParams{})
			}()
		})

		It("should set the resolver of a non pointer field", func() {
			type StructExample struct {
				Field1 CustomFieldTypeWithResolver `graphql:"field1"`
			}

			obj := gqlstruct.Struct(&StructExample{})
			Expect(obj.Fields()).To(HaveLen(1))
			Expect(obj.Fields()).To(HaveKey("field1"))
			func() {
				defer func() {
					r := recover()
					Expect(r).To(Equal("only to catch"))
				}()

				obj.Fields()["field1"].Resolve(graphql.ResolveParams{})
			}()
		})
	})
})
