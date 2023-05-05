import { Component, OnInit } from '@angular/core';
import { CategoriesService } from '../categories.service';
import CardData from '../../data/card.interface'
import Transformers from '../../utils/transformers'

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
      this.data = cats.map(c => Transformers.transformCategoryToCardData(c));
    })
  }
}
