package gqlstruct_test

import (
	"github.com/graphql-go/graphql"
	"github.com/lab259/go-graphql-struct"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"reflect"
)

var _ = Describe("ArgsOf", func() {
	type ComplexModel struct {
		FirstName  string `graphql:"!firstName"`
		MiddleName string
		LastName   string `graphql:"lastName"`
		Age        int    `graphql:"age"`
	}

	It("should generate arguments from a non pointer object", func() {
		fields, err := gqlstruct.NewEncoder().ArgsOf(reflect.TypeOf(ComplexModel{}))
		Expect(err).ToNot(HaveOccurred())
		Expect(fields).ToNot(BeNil())
		Expect(fields).To(HaveLen(3))
		Expect(fields).To(HaveKey("firstName"))
		Expect(fields["firstName"].Type).To(BeAssignableToTypeOf(graphql.NewNonNull(nil)))
		Expect(fields["firstName"].Type.(*graphql.NonNull).OfType).To(Equal(graphql.String))
		Expect(fields).To(HaveKey("lastName"))
		Expect(fields["lastName"].Type).To(Equal(graphql.String))
		Expect(fields).To(HaveKey("age"))
		Expect(fields["age"].Type).To(Equal(graphql.Int))
	})

	It("should generate arguments from a pointer", func() {
		fields, err := gqlstruct.NewEncoder().ArgsOf(reflect.TypeOf(&ComplexModel{}))
		Expect(err).ToNot(HaveOccurred())
		Expect(fields).ToNot(BeNil())
		Expect(fields).To(HaveLen(3))
		Expect(fields).To(HaveKey("firstName"))
		Expect(fields["firstName"].Type).To(BeAssignableToTypeOf(graphql.NewNonNull(nil)))
		Expect(fields["firstName"].Type.(*graphql.NonNull).OfType).To(Equal(graphql.String))
		Expect(fields).To(HaveKey("lastName"))
		Expect(fields["lastName"].Type).To(Equal(graphql.String))
		Expect(fields).To(HaveKey("age"))
		Expect(fields["age"].Type).To(Equal(graphql.Int))
	})

	It("should fail generating arguments for a non supported datatype", func() {
		_, err := gqlstruct.NewEncoder().ArgsOf(reflect.TypeOf("data"))
		Expect(err).To(HaveOccurred())
		Expect(err.Error()).To(ContainSubstring("cannot build args from a non struct"))
	})
})
