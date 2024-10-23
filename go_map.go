// Mutations

// Insert an element
m[key] = elem

// Get an element
elem = m[key]

// Delete an element
delete(m, key)

// Check if a key exists
elem, ok := m[key]

// If key is in m, then ok is true. If not, ok is false.

// If key is not in the map, then elem is the zero value for the map's element type.