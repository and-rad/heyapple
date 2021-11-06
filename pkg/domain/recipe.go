package domain

// Ingredient represents a single ingredient in a recipe.
// It's a combination of a food item and how much of that
// food is used, represented in grams.
type Ingredient struct {
	ID     uint32  `json:"id"`
	Amount float32 `json:"amount"`
}

// Recipe is a collection of ingredients.
type Recipe struct {
	repo Repository

	Items []Ingredient `json:"items"`
	ID    uint32       `json:"id"`
}

// NewRecipe returns an empty Recipe struct.
func NewRecipe(repo Repository) *Recipe {
	return &Recipe{
		repo: repo,
	}
}

// Create creates a new recipe from the current state of
// this one, assigns an id and saves it in the repository.
func (r *Recipe) Create() error {
	panic("not implemented")
}

// Load loads a recipe from the repository and stores
// the data in the current Recipe struct.
func (r *Recipe) Load(id uint32) error {
	panic("not implemented")
}

// Delete deletes this recipe from the repository and
// resets this Recipe struct.
func (r *Recipe) Delete() error {
	panic("not implemented")
}

// Add adds an ingredient to the list of ingredients. If
// the ingredient already exists, the amount is added to
// the total amount for this ingredient.
func (r *Recipe) Add(id uint32, amount float32) error {
	panic("not implemented")
}

// Remove removes an ingredient from the list of ingredients.
func (r *Recipe) Remove(id uint32) error {
	panic("not implemented")
}

// Modify changes all passed in ingredients to the new values.
func (r *Recipe) Modify(data []Ingredient) error {
	panic("not implemented")
}
