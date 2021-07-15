package main

type UserCreateInput struct{
    ID string `json:"id"`
    Password string `json:"password"`
    Name string `json:"name"`
}