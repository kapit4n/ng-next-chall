import { ComponentFixture, TestBed } from '@angular/core/testing';

import { CardMtComponent } from './card-mt.component';

describe('CardMtComponent', () => {
  let component: CardMtComponent;
  let fixture: ComponentFixture<CardMtComponent>;

  beforeEach(async () => {
    await TestBed.configureTestingModule({
      declarations: [ CardMtComponent ]
    })
    .compileComponents();

    fixture = TestBed.createComponent(CardMtComponent);
    component = fixture.componentInstance;
    fixture.detectChanges();
  });

  it('should create', () => {
    expect(component).toBeTruthy();
  });
});
