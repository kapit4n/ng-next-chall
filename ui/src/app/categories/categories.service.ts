import { HttpClient } from '@angular/common/http';
import { Injectable } from '@angular/core';

export interface Category {
  id: number;
  name: string;
} 

@Injectable({
  providedIn: 'root'
})
export class CategoriesService {
  apiURL: string = 'http://localhost:8080/categories'

  getAllCategories() {
    return this.client.get<Category[]>(this.apiURL)  
  }

  constructor(private client: HttpClient) {
    
  }


}
