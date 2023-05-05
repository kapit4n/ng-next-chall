import { Component, Input } from '@angular/core';

interface CardData {
  header: string;
  body: string;
  image: string;
}

@Component({
  selector: 'app-card-mt',
  templateUrl: './card-mt.component.html',
  styleUrls: ['./card-mt.component.css']
})
export class CardMtComponent {
  @Input('data') data: CardData;
}
