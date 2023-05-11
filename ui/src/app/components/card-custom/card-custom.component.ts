import { Component, Input, EventEmitter, Output } from '@angular/core';
import CardData from '../../data/card.interface'


@Component({
  selector: 'app-card-custom',
  templateUrl: './card-custom.component.html',
  styleUrls: ['./card-custom.component.css']
})
export class CardCustomComponent {
  @Input('data') data: CardData;

  @Output()
  redirectToShow = new EventEmitter<number>();

  onClickImage() {
    this.redirectToShow.emit(this.data.id)
    console.log("Emit redirect from card custom")
  }
  
}
