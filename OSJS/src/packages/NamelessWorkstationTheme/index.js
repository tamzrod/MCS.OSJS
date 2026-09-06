import './index.scss';
import osjs from 'osjs';
import {name} from './metadata.json';
import {register} from './src/theme.js';

osjs.register(name, register);
