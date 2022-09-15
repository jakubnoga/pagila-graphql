import { Client } from "pg";
import { ActorModel, FilmModel } from "../model";

export async function createRepository() {
  const db = new Client({
    user: "postgres",
    password: "123456",
  });

  await db.connect();

  return {
    getFilm: (title: string) => getFilm(db, title),
    getActors: (filmId: number) => getActors(db, filmId),
    getFilms: (first: number) => getFilms(db, first),
    changeLength: (title: string, newLength: number) =>
      changeLength(db, title, newLength),
  };
}

async function getFilm(db: Client, title: string): Promise<FilmModel> {
  const result = await db.query(
    `
  SELECT 
    film_id, 
    title, 
    description, 
    release_year, 
    language_id, 
    original_language_id, 
    rental_duration, 
    rental_rate, 
    length, 
    replacement_cost, 
    rating, 
    last_update, 
    special_features 
  FROM film 
  WHERE title = $1::text;`,
    [title]
  );

  return result.rows[0];
}

async function getActors(db: Client, filmId: number): Promise<ActorModel[]> {
  const result = await db.query(
    `
    select 
      a.first_name, 
      a.last_name, 
      a.actor_id 
    from 
      actor a
    join film_actor fa 
    on 
      fa.actor_id = a.actor_id 
    where fa.film_id = $1::int;
  `,
    [filmId]
  );

  return result.rows;
}

async function getFilms(db: Client, first: number): Promise<FilmModel[]> {
  const result = await db.query(
    `
    SELECT 
    film_id, 
    title, 
    description, 
    release_year, 
    language_id, 
    original_language_id, 
    rental_duration, 
    rental_rate, 
    length, 
    replacement_cost, 
    rating, 
    last_update, 
    special_features 
  FROM film 
  LIMIT $1::int`,
    [first]
  );

  return result.rows;
}

async function changeLength(
  db: Client,
  title: string,
  newLength: number
): Promise<{ film_id: number }> {
  const result = await db.query(
    `
    UPDATE film
    SET 
      length = $1::int
    WHERE
      title = $2::text
    RETURNING film_id;
  `,
    [newLength, title]
  );

  return result.rows[0];
}
