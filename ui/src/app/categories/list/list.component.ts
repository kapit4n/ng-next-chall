import { Component, OnInit } from '@angular/core';
import { CategoriesService, Category } from '../categories.service';

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.css']
})
export class ListComponent implements OnInit {

  data: Category[] = [];
  constructor(private catService: CategoriesService) {
    
  }

  ngOnInit() {
    this.catService.getAllCategories().subscribe(cats => {
      this.data = cats
    })
  }
}
