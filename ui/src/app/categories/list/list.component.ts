import { Component, OnInit } from '@angular/core';
import { CategoriesService, Category } from '../categories.service';

interface CardData {
  header: string;
  body: string;
  image: string;
}

@Component({
  selector: 'app-list',
  templateUrl: './list.component.html',
  styleUrls: ['./list.component.css']
})
export class ListComponent implements OnInit {

  data: CardData[] = [];
  constructor(private catService: CategoriesService) {

  }

  ngOnInit() {
    this.catService.getAllCategories().subscribe(cats => {
      this.data = cats.map(c => ({ header: c.name, body: c.description, image: c.image }));
    })
  }
}
