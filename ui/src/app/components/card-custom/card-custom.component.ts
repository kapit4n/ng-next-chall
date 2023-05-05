import { Component, Input } from '@angular/core';
import CardData from '../../data/card.interface'


@Component({
  selector: 'app-card-custom',
  templateUrl: './card-custom.component.html',
  styleUrls: ['./card-custom.component.css']
})
export class CardCustomComponent {
  @Input('data') data: CardData;
}
