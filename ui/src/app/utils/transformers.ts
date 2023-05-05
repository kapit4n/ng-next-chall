import CardData from "../data/card.interface";
import Category from "../data/category.interface";
import Subject from "../data/subject.interface";

export default class Transformers {
  static transformCategoryToCardData(category: Category): CardData {
    return ({ id: category.id, header: category.name, body: category.description, image: category.image })
  }

  static transformSubjectToCardData(subject: Subject): CardData {
    return ({ id: subject.id, header: subject.name, body: subject.description, image: subject.image })
  }
}