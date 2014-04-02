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