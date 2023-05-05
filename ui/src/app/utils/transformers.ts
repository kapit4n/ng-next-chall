import CardData from "../data/card.interface";
import Category from "../data/category.interface";
import Subject from "../data/subject.interface";

export default class Transformers {
  static transformCategoryToCardData(category: Category): CardData {
    return ({ header: category.name, body: category.description, image: category.image })
  }

  
  static transformSubjectToCardData(category: Subject): CardData {
    return ({ header: category.name, body: category.description, image: category.image })
  }
}