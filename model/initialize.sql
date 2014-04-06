create table post (
	id serial,
	url text not null,
	constraint post__id primary key (id)
);

create table post_meta (
	post_id integer not null,
	key text not null,
	value text not null
);

create table "user" (
	email text not null,
	password text not null,
	salt text not null,
	constraint user__email primary key (email)
);

create table user_session (
	secret text not null,
	user_email text not null,
	constraint user_session__secret primary key (secret)
);