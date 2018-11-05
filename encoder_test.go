package gqlstruct_test

import (
	"github.com/graphql-go/graphql"
	"github.com/lab259/go-graphql-struct"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
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

		obj, err := gqlstruct.NewEncoder().Struct(&StructExample{})
		Expect(err).ToNot(HaveOccurred())
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

	It("should create an object with options", func() {
		type StructExample struct {
			Field1 string `graphql:"field1"`
			Field2 string `graphql:"field2"`
		}

		obj, err := gqlstruct.NewEncoder().Field(&StructExample{}, gqlstruct.WithDescription("Description 1"))
		Expect(err).ToNot(HaveOccurred())
		Expect(obj.Description).To(Equal("Description 1"))
	})

	It("should fail when a option fails", func() {
		type StructExample struct {
			Field1 string `graphql:"field1"`
			Field2 string `graphql:"field2"`
		}

		_, err := gqlstruct.NewEncoder().Field(&StructExample{}, &erroredOption{})
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("forced error"))
	})

	It("should fail with a not recognized type", func() {
		type StructExample struct {
			Field1 []interface{} `graphql:"field1"`
		}

		_, err := gqlstruct.NewEncoder().Field(&StructExample{})
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("not recognized"))
		Expect(err.Error()).To(ContainSubstring("interface {}"))
	})

	It("should fail with a not recognized type with a pointer", func() {
		type StructExample struct {
			Field1 []*interface{} `graphql:"field1"`
		}

		_, err := gqlstruct.NewEncoder().Field(&StructExample{})
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("not recognized"))
		Expect(err.Error()).To(ContainSubstring("interface {}"))
	})
})
