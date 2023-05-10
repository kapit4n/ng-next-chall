import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';
import Category from '../data/category.interface';


@Injectable({
  providedIn: 'root'
})
export class CategoriesService {
  // move api url to config
  apiURL: string = 'http://localhost:8080/categories/'

  getAllCategories() {
    return this.client.get<Category[]>(this.apiURL)  
  }

  constructor(private client: HttpClient) {
    
  }

}
