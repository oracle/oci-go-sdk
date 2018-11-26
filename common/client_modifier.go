package common

//ClientModifier gives ability to modify a BaseClient before use.
type ClientModifier struct {
	modifiers []func(*BaseClient) (*BaseClient, error)
}

//NewClientModifier Create a new ClientModifier with optional initial modifier (may be nil).
func NewClientModifier(modifier func(*BaseClient) (*BaseClient, error)) *ClientModifier {
	clientModifier := &ClientModifier{
		modifiers: make([]func(*BaseClient) (*BaseClient, error), 0),
	}
	if modifier != nil {
		clientModifier.QueueModifier(modifier)
	}
	return clientModifier
}

//QueueModifier Queue up a new modifier
func (c *ClientModifier) QueueModifier(modifier func(*BaseClient) (*BaseClient, error)) {
	c.modifiers = append(c.modifiers, modifier)
}

//Modify the provided BaseClient with this client modifier, and return the result, or error if something
//error if something goes wrong.
func (c *ClientModifier) Modify(client *BaseClient) (*BaseClient, error) {
	if len(c.modifiers) > 0 {
		for _, modifier := range c.modifiers {
			var err error
			if client, err = modifier(client); err != nil {
				return nil, err
			}
		}
	}
	return client, nil
}
