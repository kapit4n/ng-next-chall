import { Component, Input } from '@angular/core';
interface CardData {
  header: string;
  body: string;
  image: string;
}

@Component({
  selector: 'app-card-custom',
  templateUrl: './card-custom.component.html',
  styleUrls: ['./card-custom.component.css']
})
export class CardCustomComponent {
  @Input('data') data: CardData;
}
