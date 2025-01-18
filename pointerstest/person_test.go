
package pointerstest

import (
	"testing"
)

var result Person    
var resultPtr *Person

func BenchmarkStackAllocation(b *testing.B) {
	phones := make([]string, 0, 2)
	emails := make([]string, 0, 2)
	notes := make(map[string]string)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		phones = append(phones, "123-456-7890", "098-765-4321")
		emails = append(emails, "john@example.com", "doe@example.com")
		notes["note1"] = "Some note"
		
		p := MakePerson("John", "Doe", 30)
		p.PhoneNumbers = phones
		p.Emails = emails
		p.Notes = notes
		p.Address = "123 Main St"
		
		result = p
	}
}

func BenchmarkHeapAllocation(b *testing.B) {
	phones := make([]string, 0, 2)
	emails := make([]string, 0, 2)
	notes := make(map[string]string)
	
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		phones = append(phones, "123-456-7890", "098-765-4321")
		emails = append(emails, "john@example.com", "doe@example.com")
		notes["note1"] = "Some note"
		
		p := MakePersonPointer("John", "Doe", 30)
		p.PhoneNumbers = phones
		p.Emails = emails
		p.Notes = notes
		p.Address = "123 Main St"
		
		resultPtr = p
	}
}
