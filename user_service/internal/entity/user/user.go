package user

type User struct {
    ID       int32 `json:"id" bson:"_id"`
    Username string `json:"username" bson:"username"`
    Age      int32  `json:"age" bson:"age"`
    Email    string `json:"email" bson:"email"`
    Password string `json:"password" bson:"password"`
}

type UserRequest struct {
    Username        string `json:"username" bson:"username"`
    Age             int32  `json:"age" bson:"age"`
    Password        string `json:"password" bson:"password"`
    ConfirmPassword string `json:"confirm_password" bson:"confirm_password"`
    Email           string `json:"email" bson:"email"`
}

type UserResponse struct {
    Id       int32  `json:"user_id" bson:"user_id"`
    Username string `json:"username" bson:"username"`
    Age      int32  `json:"age" bson:"age"`
    Email    string `json:"email" bson:"email"`
}

type LoginRequest struct {
    Email    string `json:"email" bson:"email"`
    Password string `json:"password" bson:"password"`
}

type LoginResponse struct {
    Token     string `json:"token" bson:"token"`
    ExpiresIn string `json:"expires_in" bson:"expires_in"`
}

type Req struct {
    Email string `json:"email" bson:"email"`
    Code  int32  `json:"code" bson:"code"`
}

type Res struct {
    Message string `json:"message" bson:"message"`
}

type GetUserRequest struct {
    ID int32 `json:"id" bson:"_id"`
}

type Empty struct{}

type ListUser struct {
    User []User `json:"user" bson:"user"`
}

type UpdateUserReq struct {
    Id       int32 `json:"user_id" bson:"user_id"`
    Username string `json:"username" bson:"username"`
    Age      int32  `json:"age" bson:"age"`
    Email    string `json:"email" bson:"email"`
}

type UpdateUserRes struct {
    Message string `json:"message" bson:"message"`
}

type UpdatePasswordReq struct {
    Id          int32 `json:"user_id" bson:"user_id"`
    OldPassword string `json:"old_password" bson:"old_password"`
    NewPassword string `json:"new_password" bson:"new_password"`
}
