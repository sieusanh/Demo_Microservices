import typeProduct from './product';

type Discount = typeProduct & {
    ratio: number   // Discount ratio
};

export default Discount;
