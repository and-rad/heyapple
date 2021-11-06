package domain

// Food represents an edible object. All nutrients are stored per 100g.
//
// TODO: add remarks about using grams for all food items internally.
type Food struct {
	repo Repository

	ID uint32 `json:"id"`

	KCal    float32 `json:"cal"`
	Fat     float32 `json:"fat"`
	Carbs   float32 `json:"carbs"`
	Protein float32 `json:"protein"`
}

// NewFood returns an empty Food struct
func NewFood(repo Repository) *Food {
	return &Food{
		repo: repo,
	}
}

// Create creates a new Food item from the current state
// of this one, assigns an id and saves it in the repository.
func (f *Food) Create() error {
	panic("not implemented")
}

// Load loads a food item from the repository and stores
// the data in the current Food struct.
func (f *Food) Load(id uint32) error {
	panic("not implemented")
}

// Delete deletes this food item from the repository and
// resets this Food struct.
func (f *Food) Delete() error {
	panic("not implemented")
}

// Modify changs properties on this food item.
func (f *Food) Modify(data map[string]interface{}) error {
	panic("not implemented")
}
