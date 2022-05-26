create database Eyebrow;

create table Sys_Setting
(
    Id int auto_increment primary key,
	SS_Key varchar(255) primary key,
	SS_Type varchar(255),
    SS_Value varchar(1500),
    SS_Dese varchar(255),
    SS_Order int
);

create table Sys_Code
(
	Id int auto_increment primary key,
	SC_Type varchar(255),
    SC_Code varchar(255),
    SC_Desc varchar(255),
    SC_Order int
);

create table Announcement
(
	Id int auto_increment primary key,
    Title varchar(255),
    Content longtext,
    Statu bool,
    Shelf_Time date,
    Create_Time date,
    IsTop bool
);

create table Banner
(
	Id int auto_increment primary key,
    ImagePath varchar(512),
    Url varchar(512),
    Statu bool,
    Banner_Order int
);

create table Reserve
(
	Id int auto_increment primary key,
    ReserveName varchar(16),
    Phone varchar(16),
    Category_Id varchar(3),
    SubCategory_Id varchar(3),
    Appointment date,
	Period bit
);

create table SubCategory
(
	Id int auto_increment primary key,
    CategoryId varchar(3),
    SubCategory_Name varchar(255),
    Price int,
    SubCategory_Order int
);

create table Portfolio
(
	Id int auto_increment primary key,
	Portfolio_Name varchar(255),
    Before_Path varchar(512),
    After_Path varchar(512),
    SubCategory_Id int,
    Portfolio_Order int,
    Statu bool
);