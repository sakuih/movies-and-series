import React, { useState, useEffect } from 'react';
import axios from 'axios'
import 'bootstrap/dist/css/bootstrap.min.css';
import './App.css';
import MovieList from './components/MovieList';
import MovieListHeading from './components/MovieListHeading';
import SearchBox from './components/SearchBox';
import AddFavourites from './components/AddFavourites';
import RemoveFavourites from './components/RemoveFavourites';

//const API_KEY = process.env.REACT_APP_APIKEY

const App = () => {
	const [movies, setMovies] = useState([])
	const [favourites, setFavourites] = useState([])
	const [searchValue, setSearchValue] = useState('')
  const [API_KEY, setAPI_KEY] = useState('')

	
	const getMovieRequest = async (searchValue) => {

		//const url = `http://www.omdbapi.com/?s=${searchValue}&apikey=${API_KEY}`

		const url = `http://localhost:3002/api/data`
		const response = axios.get(url, {
			params: { searchValue }
		})
		const responseJson = await response.json()

		if (responseJson.Search) {
			setMovies(responseJson.Search)
		}
	};


	
	useEffect(() => {
		getMovieRequest(searchValue)
	}, [searchValue])

/*
	useEffect(() => {
		getAPIKEY()
	}, [])
	useEffect(() => {
		const movieFavourites = JSON.parse(
			localStorage.getItem('react-movie-app-favourites')
		)

		setFavourites(movieFavourites)
	}, [])

	const saveToLocalStorage = (items) => {
		localStorage.setItem('react-movie-app-favourites', JSON.stringify(items))
	}
*/
	const addFavouriteMovie = (movie) => {
		const newFavouriteList = [...favourites, movie]
		setFavourites(newFavouriteList)
		//saveToLocalStorage(newFavouriteList)
	}

	const removeFavouriteMovie = (movie) => {
		const newFavouriteList = favourites.filter(
			(favourite) => favourite.imdbID !== movie.imdbID
		)

		setFavourites(newFavouriteList)
		//saveToLocalStorage(newFavouriteList)
	}

	return (
		<div className='container-fluid movie-app'>
			<div className='row d-flex align-items-center mt-4 mb-4'>
				<MovieListHeading heading='Movies' />
				<SearchBox searchValue={searchValue} setSearchValue={setSearchValue} />
        <button type="button" class="btn btn-warning m-auto">Sign In</button>
			</div>
			<div className='row'>
				<MovieList
					movies={movies}
					handleFavouritesClick={addFavouriteMovie}
					favouriteComponent={AddFavourites}
				/>
			</div>
			<div className='row d-flex align-items-center mt-4 mb-4'>
				<MovieListHeading heading='Favourites' />
			</div>
			<div className='row'>
				<MovieList
					movies={favourites}
					handleFavouritesClick={removeFavouriteMovie}
					favouriteComponent={RemoveFavourites}
				/>
			</div>
		</div>
	)
	
}

export default App;
