package generator

import "github.com/objectbox/objectbox-go/internal/generator/modelinfo"

func mergeBindingWithModelInfo(binding *Binding, modelInfo *modelinfo.ModelInfo) error {
	for _, bindingEntity := range binding.Entities {
		if modelEntity, err := getModelEntity(bindingEntity, modelInfo); err != nil {
			return err
		} else if err := mergeModelEntity(bindingEntity, modelEntity); err != nil {
			return err
		}
	}

	return nil
}

func getModelEntity(bindingEntity *Entity, modelInfo *modelinfo.ModelInfo) (*modelinfo.Entity, error) {
	if bindingEntity.Uid != 0 {
		return modelInfo.FindEntityByUid(bindingEntity.Uid)
	} else if entity, err := modelInfo.FindEntityByName(bindingEntity.Name); entity != nil {
		return entity, err
	} else {
		return modelInfo.CreateEntity()
	}
}

func mergeModelEntity(bindingEntity *Entity, modelEntity *modelinfo.Entity) (err error) {
	modelEntity.Name = bindingEntity.Name

	if bindingEntity.Id, err = modelEntity.Id.GetId(); err != nil {
		return err
	}

	if bindingEntity.Uid, err = modelEntity.Id.GetUid(); err != nil {
		return err
	}

	// add all properties from the bindings to the model and update/rename the changed ones
	for _, bindingProperty := range bindingEntity.Properties {
		if modelProperty, err := getModelProperty(bindingProperty, modelEntity); err != nil {
			return err
		} else if err := mergeModelProperty(bindingProperty, modelProperty); err != nil {
			return err
		}
	}

	// remove the missing (removed) properties
	removedProperties := make([]*modelinfo.Property, 0)
	for _, modelProperty := range modelEntity.Properties {
		if !bindingPropertyExists(modelProperty, bindingEntity) {
			removedProperties = append(removedProperties, modelProperty)
		}
	}

	for _, property := range removedProperties {
		if err := modelEntity.RemoveProperty(property); err != nil {
			return err
		}
	}

	bindingEntity.LastPropertyId = modelEntity.LastPropertyId

	return nil
}

func getModelProperty(bindingProperty *Property, modelEntity *modelinfo.Entity) (*modelinfo.Property, error) {
	if bindingProperty.Uid != 0 {
		return modelEntity.FindPropertyByUid(bindingProperty.Uid)
	} else if property, err := modelEntity.FindPropertyByName(bindingProperty.Name); property != nil {
		return property, err
	} else {
		return modelEntity.CreateProperty()
	}
}

func mergeModelProperty(bindingProperty *Property, modelProperty *modelinfo.Property) (err error) {
	modelProperty.Name = bindingProperty.Name

	if bindingProperty.Id, err = modelProperty.Id.GetId(); err != nil {
		return err
	}

	if bindingProperty.Uid, err = modelProperty.Id.GetUid(); err != nil {
		return err
	}

	return nil
}

func bindingPropertyExists(modelProperty *modelinfo.Property, bindingEntity *Entity) bool {
	for _, bindingProperty := range bindingEntity.Properties {
		if bindingProperty.Name == modelProperty.Name {
			return true
		}
	}

	return false
}