// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// User - Created user object
type User struct {
	Email     *string `json:"email,omitempty" form:"name=email"`
	FirstName *string `json:"firstName,omitempty" form:"name=firstName"`
	ID        *int64  `json:"id,omitempty" form:"name=id"`
	LastName  *string `json:"lastName,omitempty" form:"name=lastName"`
	Password  *string `json:"password,omitempty" form:"name=password"`
	Phone     *string `json:"phone,omitempty" form:"name=phone"`
	// User Status
	UserStatus *int    `json:"userStatus,omitempty" form:"name=userStatus"`
	Username   *string `json:"username,omitempty" form:"name=username"`
}

func (o *User) GetEmail() *string {
	if o == nil {
		return nil
	}
	return o.Email
}

func (o *User) GetFirstName() *string {
	if o == nil {
		return nil
	}
	return o.FirstName
}

func (o *User) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *User) GetLastName() *string {
	if o == nil {
		return nil
	}
	return o.LastName
}

func (o *User) GetPassword() *string {
	if o == nil {
		return nil
	}
	return o.Password
}

func (o *User) GetPhone() *string {
	if o == nil {
		return nil
	}
	return o.Phone
}

func (o *User) GetUserStatus() *int {
	if o == nil {
		return nil
	}
	return o.UserStatus
}

func (o *User) GetUsername() *string {
	if o == nil {
		return nil
	}
	return o.Username
}
